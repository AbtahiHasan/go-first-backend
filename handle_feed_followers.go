package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{} 

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newFeedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		FeedID: params.FeedId,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	responseFormatter(w, 201, newFeedFollow)

}