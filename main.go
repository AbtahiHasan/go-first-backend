package main

import (
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
	router.Mount("/v1",v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":"+port,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on port", port)
}