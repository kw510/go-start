package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kw510/pg-go-react-stack/pkg/datastore"
	"github.com/kw510/pg-go-react-stack/pkg/server"
)

func main() {
	// Connect to datastore
	if err := datastore.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer datastore.Shutdown()

	// Define our routing engine
	r := gin.Default()

	// Middleware to handle cors requests
	r.Use(cors.Default())

	// Setup routing
	server.Init(r)

	// Start the server
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
}
