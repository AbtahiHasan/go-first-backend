package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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

		
	apiCfg := apiConfig{
		DB: database.New(connection),
	}

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
	v1Router.Get("/users", apiCfg.handleGetUser)
	router.Mount("/v1",v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":"+port,
	}
	
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on port", port)
}