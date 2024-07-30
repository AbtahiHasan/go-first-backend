package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/auth"
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
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		errorFormatter(w,403, fmt.Sprintf("auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(),apiKey)
	if err != nil {
		errorFormatter(w,404, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}
	responseFormatter(w,200,databaseUserToUser(user))
}