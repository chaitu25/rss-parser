package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type FeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RssFeed struct {
	Channel struct {
		Title       string     `xml:"title"`
		Description string     `xml:"description"`
		Link        string     `xml:"link"`
		Language    string     `xml:"language"`
		Items       []FeedItem `xml:"item"`
	} `xml:"channel"`
}

func urlToFeed(url string) (RssFeed, error) {
	feed := RssFeed{}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := client.Get(url)
	if err != nil {
		return feed, err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return feed, err
	}
	return feed, nil
}
