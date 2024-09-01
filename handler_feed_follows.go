package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/flamme97/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)


func(apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: 			 uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: 	 user.ID,
		FeedID: 	 params.FeedID,

	})
	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("can't load feeds: %v", err))
		return

	}

	respondWithJSON(w, 201, databseFeedFollowToFeedFollow(feedFollow))
}




func(apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID);
	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("couldn't get follows: %v", err))
		return

	}

	respondWithJSON(w, 201, databaseFeedsFollowsToFeedFollows(feedFollows))
}

func(apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")

	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't unfollow : %v", err))
		return 
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't unfollow : %v", err))
		return 
	}
	respondWithJSON(w, 200, struct{}{})

}


	