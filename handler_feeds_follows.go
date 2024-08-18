package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mojtabafarzaneh/rssagg/internal/database"
)

func (apiconf *apiConfig) CreateFeedsFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	param := Params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("couldn't decode data %s", err))
		return
	}
	feedFollows, err := apiconf.DB.CreateFeedsfollows(r.Context(), database.CreateFeedsfollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    param.FeedID,
	})
	if err != nil {
		responseWithErr(w, 404, fmt.Sprintf("couldn't find the requested user or feed %s", err))
		return
	}

	responseWithJSON(w, 201, DatabaseFeedsFollowsToFeedsFollows(feedFollows))
}

func (apiconf *apiConfig) GetFeedsfollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := apiconf.DB.GetFeedsFollos(r.Context(), user.ID)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("couldnt get the feed follows: %s", err))
	}

	responseWithJSON(w, 200, SliceFeedsFollows(feedsFollows))
}

func (apiconf *apiConfig) DeleteFeedsFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	idStr := chi.URLParam(r, "feedfollowid")
	uuid, err := uuid.Parse(idStr)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("couldnt parse feed follows uuid: %s", err))

	}

	err = apiconf.DB.DeleteFeedfollows(r.Context(), database.DeleteFeedfollowsParams{
		ID:     uuid,
		UserID: user.ID,
	})
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("couldnt delete the desired feed: %s", err))

	}

	responseWithJSON(w, 204, struct{}{})
}
