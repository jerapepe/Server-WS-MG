package main

import (
	"Server-WS/pkg/mongo"
	"Server-WS/pkg/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	mongoClient := mongo.GetMongoClient()
	defer mongoClient.Disconnect(context.Background())

	router := mux.NewRouter()
	routes.SetRoutes(router)
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server 8000")
	log.Fatal(srv.ListenAndServe())
}
