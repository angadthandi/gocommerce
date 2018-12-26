package main

import (
	"net/http"

	"github.com/angadthandi/gocommerce/config"
	log "github.com/angadthandi/gocommerce/log"
	"github.com/angadthandi/gocommerce/route"

	// // https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	log.Infof("Started main: %v", "goapp")

	// dbUrl := "mongodb://" +
	// 	config.MongoDBUsername + ":" +
	// 	config.MongoDBPassword + "@" +
	// 	config.MongoDBServiceName +
	// 	config.MongoDBPort
	// dbClient, err := mongo.Connect(context.Background(), dbUrl, nil)
	// if err != nil {
	// 	log.Fatalf("mongodb connection error : %v", err)
	// }

	// log.Debug("Connected to mongodb golangapp database")
	// defer dbClient.Disconnect(context.Background())

	// dbRef := dbClient.Database("GolangappDB")
	// log.Debug("Initialized mongodb golangapp database")
	var dbRef *mongo.Database

	route.Handle(dbRef)

	log.Printf("Listening on Port: %v", config.ServerPort)
	// start http web server
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}
