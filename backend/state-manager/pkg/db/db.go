package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("books")
	// add a new movie "Doraemon"

	// delete all
	_, err = coll.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	sampleBooks := []interface{}{
		Book{"Doraemon", "Fujiko F. Fujio"},
		Book{"Harry Potter", "J. K. Rowling"},
		Book{"The Lord of the Rings", "J. R. R. Tolkien"},
		Book{"The Hobbit", "J. R. R. Tolkien"},
		Book{"And Then There Were None", "Agatha Christie"},
		Book{"Alice's Adventures in Wonderland", "Lewis Carroll"},
		Book{"Dream of the Red Chamber", "Cao Xueqin"},
		Book{"The Lion, the Witch and the Wardrobe", "C. S. Lewis"},
		Book{"She: A History of Adventure", "H. Rider Haggard"},
	}

	_, err = coll.InsertMany(context.Background(), sampleBooks)
	if err != nil {
		panic(err)
	}

	filter := bson.D{{Key: "author", Value: "J. R. R. Tolkien"}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Book
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	log.Println(results)
}
