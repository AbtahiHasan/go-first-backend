package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not found")
	}


	router := chi.NewRouter()

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