package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/chaitu25/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func startScrapping(db *database.Queries, threads int, interval time.Duration) error {
	log.Printf("Starting scraper with %d threads", threads)

	t := time.NewTicker(interval)
	for ; ; <-t.C {
		feeds, _ := db.GetNextFeedToFetch(context.Background(), int32(threads))
		wg := &sync.WaitGroup{}
		for _, v := range feeds {
			wg.Add(1)
			go scrape(db, wg, v)
		}
		wg.Wait()

	}
}

func scrape(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFieldAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error while updating feed fetch status: %s ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error while fetching feed: %s", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		descrip := sql.NullString{}
		if item.Description != "" {
			descrip.String = item.Description
			descrip.Valid = true
		} else {
			descrip.Valid = false
		}
		urlstr := sql.NullString{
			String: item.Link,
			Valid:  true,
		}
		uidString := uuid.NullUUID{
			UUID:  feed.ID,
			Valid: true,
		}
		pubAt, _ := time.Parse(time.RFC1123Z, item.PubDate)
		db.CreatePosts(context.Background(), database.CreatePostsParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Tile:        item.Title,
			Description: descrip,
			PublishedAt: pubAt,
			Url:         urlstr,
			FeedID:      uidString,
		})
	}

	log.Printf("Feed %s is collected, total items found: %v", feed.Name, len(rssFeed.Channel.Items))

}
