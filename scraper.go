package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scraper with concurrency %d and time between requests %s", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _,feed := range feeds {
			wg.Add(1)
			go scrapeFeed(*db,wg, feed)
		}
		wg.Wait()
	}
}

func  scrapeFeed(db database.Queries,wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error parsing feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title)
	}
	log.Printf("feed %s collected, %v post found", feed.Name, len(rssFeed.Channel.Item))
}