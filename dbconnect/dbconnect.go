package dbconnect

import (
	log "github.com/angadthandi/gocommerce/log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func Conn() *mongo.Database {
	var dbRef *mongo.Database
	log.Print("TODO DB connection!")

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

	return dbRef
}
