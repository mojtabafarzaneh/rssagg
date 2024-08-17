package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mojtabafarzaneh/rssagg/internal/database"
)

func (apiconf *apiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type Params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	param := Params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("err parsing json: %s", err))
		return
	}

	feed, err := apiconf.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      param.Name,
		Url:       param.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithErr(w, 400, fmt.Sprint("couldn't create user: ", err))
		return
	}

	responseWithJSON(w, 201, DatabaseFeedToFeed(feed))
}

func (apiconf *apiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiconf.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithErr(w, 404, fmt.Sprintf("didn't find any feeds %s", err.Error()))
	}

	responseWithJSON(w, 200, DatabaseFeedsToFeeds(feeds))
}
