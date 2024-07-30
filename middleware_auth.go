package main

import (
	"fmt"
	"net/http"

	"github.com/AbtahiHasan/go-first-backend/internal/auth"
	"github.com/AbtahiHasan/go-first-backend/internal/database"
)

type authHandler func(http.ResponseWriter,*http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	
		handler(w,r,user)
	}
}