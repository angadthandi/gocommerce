package dbconnect

import (
	"context"

	"github.com/angadthandi/gocommerce/config"
	log "github.com/angadthandi/gocommerce/log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

func Conn() *mongo.Client {
	var dbClient *mongo.Client

	dbUrl := "mongodb://" +
		config.MongoDBUsername + ":" +
		config.MongoDBPassword + "@" +
		config.MongoDBServiceName +
		config.MongoDBPort
	dbClient, err := mongo.Connect(context.Background(), dbUrl, nil)
	if err != nil {
		log.Fatalf("mongodb connection error : %v", err)
	}

	err = dbClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("mongodb Ping error : %v", err)
	}

	log.Debug("Connected to mongodb gocommerce client")

	return dbClient
}
