/*
	DB Controller
*/

package db

import (
	"context"
	"fmt"
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

func (db *DBHandler) Close(ctx context.Context) {
	slog.Default().Debug("Closing connection to DB...")
	if err := db.stateManagerDB.Client().Disconnect(ctx); err != nil {
		slog.Default().Error("DB Connection Closing failed")
		return
	}
	slog.Default().Debug("DB Connection is successfully closed")
}

/*
	Point
*/

func (db *DBHandler) UpdatePoint(ctx context.Context, PointAndState *statev1.PointAndState) error {
	collection := db.stateManagerDB.Collection("points")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id": PointAndState.Id},
		bson.M{"$set": bson.M{"state": PointAndState.State}},
	)
	return fmt.Errorf("update point failed `%w`", err)
}

func (db *DBHandler) AddPoint(ctx context.Context, PointAndState *statev1.PointAndState) error {
	collection := db.stateManagerDB.Collection("points")
	_, err := collection.InsertOne(ctx, PointAndState)
	if err != nil {
		return fmt.Errorf("insert point failed `%w`", err)
	}
	return nil
}

func (db *DBHandler) GetPoint(ctx context.Context, pointId string) (*statev1.PointAndState, error) {
	collection := db.stateManagerDB.Collection("points")
	var result *statev1.PointAndState
	err := collection.FindOne(ctx, bson.M{"id": pointId}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("get point failed `%w`", err)
	}
	return result, nil
}

func (db *DBHandler) GetPoints(ctx context.Context) ([]*statev1.PointAndState, error) {
	collection := db.stateManagerDB.Collection("points")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		slog.Default().Warn("Get Points failed", slog.Any("err", err))
		return nil, fmt.Errorf("get points failed `%w`", err)
	}
	var result []*statev1.PointAndState
	if err := cursor.All(ctx, &result); err != nil {
		slog.Default().Warn("Get Points failed", slog.Any("err", err))
		return nil, fmt.Errorf("get points failed `%w`", err)
	}
	return result, nil
}

/*
	Stop
*/

func (db *DBHandler) UpdateStop(ctx context.Context, stop *statev1.StopAndState) error {
	collection := db.stateManagerDB.Collection("stops")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id": stop.Id},
		bson.M{"$set": bson.M{"state": stop.State}},
	)

	if err != nil {
		return fmt.Errorf("update stop failed `%w`", err)
	}
	return nil
}

func (db *DBHandler) AddStop(ctx context.Context, stop *statev1.StopAndState) error {
	collection := db.stateManagerDB.Collection("stops")
	_, err := collection.InsertOne(ctx, stop)
	if err != nil {
		return fmt.Errorf("insert stop failed `%w`", err)
	}
	return nil
}

func (db *DBHandler) GetStop(ctx context.Context, stopId string) (*statev1.StopAndState, error) {
	collection := db.stateManagerDB.Collection("stops")
	var result *statev1.StopAndState
	err := collection.FindOne(ctx, bson.M{"id": stopId}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("get stop failed `%w`", err)
	}
	return result, nil
}

func (db *DBHandler) GetStops(ctx context.Context) ([]*statev1.StopAndState, error) {
	collection := db.stateManagerDB.Collection("stops")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("get stops failed `%w`", err)
	}
	var result []*statev1.StopAndState
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("get stops failed `%w`", err)
	}
	return result, nil
}

/*
	Block
*/

func (db *DBHandler) AddBlock(ctx context.Context, block *statev1.BlockState) error {
	collection := db.stateManagerDB.Collection("blocks")
	_, err := collection.InsertOne(ctx, block)
	if err != nil {
		return fmt.Errorf("insert block failed `%w`", err)
	}
	return nil
}

func (db *DBHandler) UpdateBlock(ctx context.Context, block *statev1.BlockState) error {
	collection := db.stateManagerDB.Collection("blocks")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"blockid": block.BlockId},
		bson.M{"$set": bson.M{"state": block.State}},
	)
	if err != nil {
		return fmt.Errorf("update block failed `%w`", err)
	}
	return nil
}

func (db *DBHandler) GetBlock(ctx context.Context, blockId string) (*statev1.BlockState, error) {
	collection := db.stateManagerDB.Collection("blocks")
	var result *statev1.BlockState
	err := collection.FindOne(ctx, bson.M{"blockid": blockId}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("get block failed `%w`", err)
	}
	return result, nil
}

func (db *DBHandler) GetBlocks(ctx context.Context) ([]*statev1.BlockState, error) {
	collection := db.stateManagerDB.Collection("blocks")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("get blocks failed `%w`", err)
	}
	var result []*statev1.BlockState
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("get blocks failed `%w`", err)
	}
	return result, nil
}
