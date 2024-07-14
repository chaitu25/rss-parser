package main

import (
	"time"

	"github.com/chaitu25/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	API_KEY   string    `json:"api_key"`
}

type Feed struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
}

type FeedFollower struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	FeedID    string    `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tile        string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         *string   `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID.String(),
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		API_KEY:   dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID.String(),
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID.String(),
	}
}

func databaseFeedFollowerToFeedFollower(dbFeedFollower database.FeedFollower) FeedFollower {
	return FeedFollower{
		ID:        dbFeedFollower.ID.String(),
		CreatedAt: dbFeedFollower.CreatedAt,
		UpdatedAt: dbFeedFollower.UpdatedAt,
		UserID:    dbFeedFollower.UserID.String(),
		FeedID:    dbFeedFollower.FeedID.String(),
	}
}

func databaseFeedFollowersToFeedFollowers(dbFeedFollower []database.FeedFollower) []FeedFollower {
	followers := make([]FeedFollower, 0)
	for _, v := range dbFeedFollower {

		followers = append(followers, databaseFeedFollowerToFeedFollower(v))
	}
	return followers
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Tile:        post.Tile,
		Description: &post.Description.String,
		PublishedAt: post.PublishedAt,
		Url:         &post.Url.String,
		FeedID:      post.FeedID.UUID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	followers := make([]Post, 0)
	for _, v := range posts {

		followers = append(followers, databasePostToPost(v))
	}
	return followers
}
