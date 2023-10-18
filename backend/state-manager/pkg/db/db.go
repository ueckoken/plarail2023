/*
	DB Controller
*/

package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb connection
var client *mongo.Client

type Book struct {
	Title  string
	Author string
}

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

func close() {
	log.Println("Closing connection to DB...")
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	log.Println("Connection closed!")
}

func addBook(book Book) error {
	collection := client.Database("test").Collection("books")
	_, err := collection.InsertOne(context.TODO(), book)
	return err
}

func getBooks() ([]Book, error) {
	collection := client.Database("test").Collection("books")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var books []Book
	for cursor.Next(context.TODO()) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func getBook(title string) (Book, error) {
	var book Book
	collection := client.Database("test").Collection("books")
	err := collection.FindOne(context.TODO(), bson.M{
		"title": title,
	}).Decode(&book)
	return book, err
}

func deleteBook(book Book) error {
	collection := client.Database("test").Collection("books")
	_, err := collection.DeleteOne(context.TODO(), bson.M{
		"title":  book.Title,
		"author": book.Author,
	})
	return err
}
