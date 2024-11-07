package tools

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RenameField(collection *mongo.Collection, oldFieldName string, newFieldName string) error {
	_, err := collection.UpdateMany(context.Background(), bson.M{}, bson.M{"$rename": bson.M{oldFieldName: newFieldName}})
	if err != nil {
		return err
	}
	return nil
}

func RemoveField(collection *mongo.Collection, fieldName string) error {
	_, err := collection.UpdateMany(context.Background(), bson.M{}, bson.M{"$unset": bson.M{fieldName: ""}})
	if err != nil {
		return err
	}
	return nil
}
