package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbtahiHasan/go-first-backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}


func main() {

	
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not found")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == ""{
		log.Fatal("DB_URL not found")
	}

	connection, err :=sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("DB connection failed")
	}

	db := database.New(connection)
	apiCfg := apiConfig{
		DB: db,
	}
	go startScraping(db,10, time.Minute)
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: 	[]string{"https://*","http://*"},
		AllowedMethods: 	[]string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: 	[]string{"*"},
		ExposedHeaders: 	[]string{"Link"},
		AllowCredentials: 	false,
		MaxAge: 			300,
	}))


	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)
	v1Router.Post("/follow-feed", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/follow-feed", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollows))
	v1Router.Delete("/un-follow-feed/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))
	router.Mount("/v1",v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":"+port,
	}
	
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on port", port)
}