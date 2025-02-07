package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`	
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newUser, userCreatingErr := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if userCreatingErr != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error creating user: %v", userCreatingErr))
		return
	}

	responseFormatter(w, 201, databaseUserToUser(newUser))
}
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	
	responseFormatter(w,200,databaseUserToUser(user))
}
func (apiCfg *apiConfig) handleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		errorFormatter(w, 400, fmt.Sprintf("Error getting posts: %v", err))
		return
	}

	responseFormatter(w,200,posts)
}

