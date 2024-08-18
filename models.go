package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mojtabafarzaneh/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseUserToUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKey:    dbuser.ApiKey,
	}
}

func DatabaseFeedToFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}
	return feeds
}

func DatabaseFeedsFollowsToFeedsFollows(dbFeedsFollows database.FeedsFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedsFollows.ID,
		CreatedAt: dbFeedsFollows.CreatedAt,
		UpdatedAt: dbFeedsFollows.UpdatedAt,
		UserID:    dbFeedsFollows.UserID,
		FeedID:    dbFeedsFollows.FeedID,
	}
}

func SliceFeedsFollows(dbFeedsFollows []database.FeedsFollow) []FeedFollows {
	feeds := []FeedFollows{}

	for _, dbFeedFollow := range dbFeedsFollows {
		feeds = append(feeds, DatabaseFeedsFollowsToFeedsFollows(dbFeedFollow))
	}
	return feeds
}
