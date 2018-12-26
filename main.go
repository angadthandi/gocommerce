package main

import (
	"net/http"

	"github.com/angadthandi/gocommerce/config"
	"github.com/angadthandi/gocommerce/dbconnect"
	log "github.com/angadthandi/gocommerce/log"
	"github.com/angadthandi/gocommerce/route"
	// // https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
)

func main() {
	dbRef := dbconnect.Conn()

	route.Handle(dbRef)

	log.Printf("Listening on Port: %v", config.ServerPort)
	// start http web server
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}
