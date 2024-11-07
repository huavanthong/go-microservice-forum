package tools

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CopyData(source *mongo.Collection, dest *mongo.Collection) error {
	cursor, err := source.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	var documents []interface{}
	for cursor.Next(context.Background()) {
		var doc bson.M
		err := cursor.Decode(&doc)
		if err != nil {
			return err
		}
		documents = append(documents, doc)
	}
	_, err = dest.InsertMany(context.Background(), documents)
	if err != nil {
		return err
	}
	return nil
}
