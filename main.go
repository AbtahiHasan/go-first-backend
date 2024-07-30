package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not found")
	}


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	server := &http.Server{
		Handler: router,
		Addr: ":"+port,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on port", port)
}