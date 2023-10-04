package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PrithviPy/go-automation-testing/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const DBName = "go_buddies_db"

var MonGoClient *mongo.Client

var DbContext *context.Context

func Close(cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := MonGoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Print("***Connection Closed***")
}

func Connect(uri string) (context.CancelFunc, error) {
	log.Println("----Trying to conencto to DB----")
	log.Printf("%v", uri)
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	DbContext = &ctx
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	MonGoClient = client
	return cancel, err
}

func Ping() error {
	if err := MonGoClient.Ping(*DbContext, readpref.Primary()); err != nil {
		return err
	}
	log.Println("----connected successfully----")
	return nil
}

func InsertOne(collectionName string, doc interface{}) (interface{}, error) {
	collection := *MonGoClient.Database(DBName).Collection(collectionName)
	result, err := collection.InsertOne(*DbContext, doc)
	if result == nil {
		log.Fatal("Something not roight while inserting")
	}
	return result, err
}

func FindOne(collectionName string, filterDoc interface{}, bindingInterface interface{}) (interface{}, error) {
	collection := *MonGoClient.Database(DBName).Collection(collectionName)
	filter, _ := utils.CreateBSONWithNonEmptyFields(filterDoc)
	fmt.Printf("Trying to find %v", filter)
	err := collection.FindOne(*DbContext, filter).Decode(bindingInterface)
	return bindingInterface, err
}

func UpdateOne(collectionName string, doc interface{}) (interface{}, error) {
	collection := *MonGoClient.Database(DBName).Collection(collectionName)
	filter, _ := utils.CreateBSONWithNonEmptyFields(doc)
	fmt.Printf("Trying to find %v", filter)
	_, err := collection.UpdateOne(*DbContext, filter, doc)
	return doc, err
}

func DeleteOne(collectionName string, doc interface{}) (interface{}, error) {
	collection := *MonGoClient.Database(DBName).Collection(collectionName)
	filter, _ := utils.CreateBSONWithNonEmptyFields(doc)
	fmt.Printf("Trying to find %v", filter)
	_, err := collection.DeleteOne(*DbContext, filter)
	return doc, err
}
