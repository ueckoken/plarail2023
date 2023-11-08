/*
	DB Controller
*/

package db

import (
	"context"
	"log"
	"os"

	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb connection
var client *mongo.Client

func Open() {
	var err error
	//log.Println("Connecting to MongoDB...")
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No MONGODB_URI set")
	}
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	//log.Println("Connected to DB!")
}

func C() {
	//log.Println("Closing connection to DB...")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Println(err)
	}
	//log.Println("Connection closed!")
}

/*
	Point
*/

func UpdatePoint(PointAndState *statev1.PointAndState) error {
	collection := client.Database("state-manager").Collection("points")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": PointAndState.Id},
		bson.M{"$set": bson.M{"state": PointAndState.State}},
	)
	if err != nil {
		return err
	}
	return nil
}

func AddPoint(PointAndState *statev1.PointAndState) error {
	collection := client.Database("state-manager").Collection("points")
	_, err := collection.InsertOne(context.Background(), PointAndState)
	if err != nil {
		return err
	}
	return nil
}

func GetPoint(pointId string) (*statev1.PointAndState, error) {
	collection := client.Database("state-manager").Collection("points")
	var result *statev1.PointAndState
	err := collection.FindOne(context.Background(), bson.M{"id": pointId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetPoints() []*statev1.PointAndState {
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

/*
	Stop
*/

func UpdateStop(stop *statev1.StopAndState) error {
	collection := client.Database("state-manager").Collection("stops")

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": stop.Id},
		bson.M{"$set": bson.M{"state": stop.State}},
	)

	if err != nil {
		return err
	}
	return nil
}

func AddStop(stop *statev1.StopAndState) error {
	collection := client.Database("state-manager").Collection("stops")
	_, err := collection.InsertOne(context.Background(), stop)
	if err != nil {
		return err
	}
	return nil
}

func GetStop(stopId string) (*statev1.StopAndState, error) {
	collection := client.Database("state-manager").Collection("stops")
	var result *statev1.StopAndState
	err := collection.FindOne(context.Background(), bson.M{"id": stopId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetStops() []*statev1.StopAndState {
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

/*
	Block
*/

func AddBlock(block *statev1.BlockState) error {
	collection := client.Database("state-manager").Collection("blocks")
	_, err := collection.InsertOne(context.Background(), block)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlock(block *statev1.BlockState) error {
	collection := client.Database("state-manager").Collection("blocks")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"blockid": block.BlockId},
		bson.M{"$set": bson.M{"state": block.State}},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetBlock(blockId string) (*statev1.BlockState, error) {
	collection := client.Database("state-manager").Collection("blocks")
	var result *statev1.BlockState
	err := collection.FindOne(context.Background(), bson.M{"blockid": blockId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBlocks() ([]*statev1.BlockState, error) {
	collection := client.Database("state-manager").Collection("blocks")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var result []*statev1.BlockState
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}
