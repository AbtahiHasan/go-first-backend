package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/database"
	"github.com/go-chi/chi"
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
func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	
	newFeedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("getting feed follow error: %v", err))
		return
	}

	responseFormatter(w, 200, newFeedFollow)

}
func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	
	feedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error parsing follow id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:feedFollowId, 
		UserID:user.ID,
	})

	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error deleting follow: %v", err))
		return
	}
	responseFormatter(w, 200, struct{}{})

}