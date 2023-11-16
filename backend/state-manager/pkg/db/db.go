/*
	DB Controller
*/

package db

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	stateManagerDB *mongo.Database
}

func Open(ctx context.Context, opts *options.ClientOptions) (*DBHandler, error) {
	var err error
	slog.Default().Debug("Connecting to MongoDB...")
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		slog.Default().Error("database connection failed", slog.Any("err", err))
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		slog.Error("DB ping failed", slog.Any("err", err))
		return nil, fmt.Errorf("DB Ping failed `%w`", err)
	}
	slog.Default().Debug("connected to DB")
	return &DBHandler{
		stateManagerDB: client.Database("state-manager"),
	}, nil
}

func (db *DBHandler) Close() {
	slog.Default().Debug("Closing connection to DB...")
	// TODO: contextを受けて、その子contextをDBクライアントに渡す
	if err := db.stateManagerDB.Client().Disconnect(context.TODO()); err != nil {
		slog.Default().Error("DB Connection Closing failed")
		log.Println(err)
	}
	slog.Default().Debug("DB Connection is successfully closed")
}

/*
	Point
*/

func (db *DBHandler) UpdatePoint(PointAndState *statev1.PointAndState) error {
	collection := db.stateManagerDB.Collection("points")
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

func (db *DBHandler) AddPoint(PointAndState *statev1.PointAndState) error {
	collection := db.stateManagerDB.Collection("points")
	_, err := collection.InsertOne(context.Background(), PointAndState)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBHandler) GetPoint(pointId string) (*statev1.PointAndState, error) {
	collection := db.stateManagerDB.Collection("points")
	var result *statev1.PointAndState
	err := collection.FindOne(context.Background(), bson.M{"id": pointId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DBHandler) GetPoints() []*statev1.PointAndState {
	collection := db.stateManagerDB.Collection("points")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		slog.Default().Warn("Get Points failed", slog.Any("err", err))
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

func (db *DBHandler) UpdateStop(stop *statev1.StopAndState) error {
	collection := db.stateManagerDB.Collection("stops")

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

func (db *DBHandler) AddStop(stop *statev1.StopAndState) error {
	collection := db.stateManagerDB.Collection("stops")
	_, err := collection.InsertOne(context.Background(), stop)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBHandler) GetStop(stopId string) (*statev1.StopAndState, error) {
	collection := db.stateManagerDB.Collection("stops")
	var result *statev1.StopAndState
	err := collection.FindOne(context.Background(), bson.M{"id": stopId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DBHandler) GetStops() []*statev1.StopAndState {
	collection := db.stateManagerDB.Collection("stops")
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

func (db *DBHandler) AddBlock(block *statev1.BlockState) error {
	collection := db.stateManagerDB.Collection("blocks")
	_, err := collection.InsertOne(context.Background(), block)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBHandler) UpdateBlock(block *statev1.BlockState) error {
	collection := db.stateManagerDB.Collection("blocks")
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

func (db *DBHandler) GetBlock(blockId string) (*statev1.BlockState, error) {
	collection := db.stateManagerDB.Collection("blocks")
	var result *statev1.BlockState
	err := collection.FindOne(context.Background(), bson.M{"blockid": blockId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DBHandler) GetBlocks() ([]*statev1.BlockState, error) {
	collection := db.stateManagerDB.Collection("blocks")
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
