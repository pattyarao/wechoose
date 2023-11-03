package database

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MG MongoInstance

func Connect() error {
	godotenv.Load()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		return err
	}

	db := client.Database(os.Getenv("DB_NAME"))

	if err != nil {
		return err
	}

	MG = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func Disconnect() error {
	defer func() {
		if err := MG.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return nil
}

func CreateUniqueIndex() error {
	collection := MG.Db.Collection("users")

	usernameIndex := mongo.IndexModel{
		Keys:    bson.M{"user_name": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), usernameIndex)
	if err != nil {
		panic(err)
	}
	return nil
}
