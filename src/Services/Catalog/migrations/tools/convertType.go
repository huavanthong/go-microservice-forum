func ConvertFieldType(collection *mongo.Collection, fieldName string, fromType bson.Type, toType bson.Type) error {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var doc bson.M
		err := cursor.Decode(&doc)
		if err != nil {
			return err
		}
		if value, ok := doc[fieldName]; ok && value != nil {
			if bsonType := value.(primitive.Type); bsonType == fromType {
				switch toType {
				case bson.TypeInt32:
					doc[fieldName] = int32(value.(int))
				case bson.TypeInt64:
					doc[fieldName] = int64(value.(int))
				case bson.TypeDouble:
					doc[fieldName] = float64(value.(int))
				case bson.TypeString:
					doc[fieldName] = strconv.Itoa(value.(int))
				}
				_, err := collection.UpdateOne(context.Background(), bson.M{"_id": doc["_id"]}, bson.M{"$set": doc})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
