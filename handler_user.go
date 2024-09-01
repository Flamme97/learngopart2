package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/flamme97/rssagg/internal/database"
	"github.com/google/uuid"
)


func(apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: 			 uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: 		 params.Name,

	})
	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("can't create user: %v", err))
		return

	}

	respondWithJSON(w, 201, dataUserToUser(user))
}

func(apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	respondWithJSON(w, 200, dataUserToUser(user))
}

func(apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User){
	posts, err := apiCfg.DB.GetPOstsForUser(r.Context(), database.GetPOstsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil{
		respondWithJSON(w, 400,fmt.Sprintf("couldn't get posts: %v", err))
		return
	}
	respondWithJSON(w, 200, databasepostsToPost(posts))

}