package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name 	string 	`json:"name"`	
		Url		string 	`json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newFeed, feedCreatingErr := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if feedCreatingErr != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error creating feed: %v", feedCreatingErr))
		return
	}

	responseFormatter(w, 201, newFeed)
}


func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}

	responseFormatter(w, 200, feeds)
}