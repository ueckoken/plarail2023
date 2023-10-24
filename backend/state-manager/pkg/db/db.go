/*
	DB Controller
*/

package main

import (
	"context"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb connection
var client *mongo.Client

func open() {
	var err error
	log.Println("Connecting to MongoDB...")
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No MONGODB_URI set")
	}
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB!")
}

func c() {
	log.Println("Closing connection to DB...")
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	log.Println("Connection closed!")
}

func setPoint(PointAndState *statev1.PointAndState) {
	collection := client.Database("state-manager").Collection("points")
	_, err := collection.InsertOne(context.Background(), PointAndState)
	if err != nil {
		panic(err)
	}
}

func getPoint(pointId string) *statev1.PointAndState {
	collection := client.Database("state-manager").Collection("points")
	var result *statev1.PointAndState
	err := collection.FindOne(context.Background(), bson.M{"id": pointId}).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func getPoints() []*statev1.PointAndState {
	collection := client.Database("state-manager").Collection("points")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	var result []*statev1.PointAndState
	if err = cursor.All(context.Background(), &result); err != nil {
		panic(err)
	}
	return result
}

func setStop(stop *statev1.StopAndState) {
	collection := client.Database("state-manager").Collection("stops")
	_, err := collection.InsertOne(context.Background(), stop)
	if err != nil {
		panic(err)
	}
}

func getStop(stopId string) *statev1.StopAndState {
	collection := client.Database("state-manager").Collection("stops")
	var result *statev1.StopAndState
	err := collection.FindOne(context.Background(), bson.M{"id": stopId}).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func getStops() []*statev1.StopAndState {
	collection := client.Database("state-manager").Collection("stops")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	var result []*statev1.StopAndState
	if err = cursor.All(context.Background(), &result); err != nil {
		panic(err)
	}
	return result
}
