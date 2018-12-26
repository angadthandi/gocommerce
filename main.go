package main

import (
	"net/http"

	"github.com/angadthandi/gocommerce/config"
	"github.com/angadthandi/gocommerce/dbconnect"
	"github.com/angadthandi/gocommerce/gosocket"
	log "github.com/angadthandi/gocommerce/log"
	"github.com/angadthandi/gocommerce/route"
	// // https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
)

func main() {
	dbRef := dbconnect.Conn()

	// start hub
	// for creating websocket conns
	hub := gosocket.NewHub()
	go hub.Run()

	route.Handle(dbRef, hub)

	log.Printf("Listening on Port: %v", config.ServerPort)
	// start http web server
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}
