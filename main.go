package main

import (
	"context"
	"net/http"

	"github.com/angadthandi/gocommerce/config"
	"github.com/angadthandi/gocommerce/dbconnect"
	"github.com/angadthandi/gocommerce/gosocket"
	log "github.com/angadthandi/gocommerce/log"
	"github.com/angadthandi/gocommerce/registry"
	"github.com/angadthandi/gocommerce/route"
	// // https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
)

func main() {
	dbClient := dbconnect.Conn()
	defer dbClient.Disconnect(context.Background())
	dbRef := dbClient.Database("gocommercedb")
	log.Debug("Initialized mongodb gocommerce database")

	clientRegistry := registry.NewRegistry()

	// start hub
	// for creating websocket conns
	hub := gosocket.NewHub()
	go hub.Run(clientRegistry)

	route.Handle(dbRef, hub, clientRegistry)

	log.Printf("Listening on Port: %v", config.ServerPort)
	// start http web server
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}
