package mongo

import (
	"context"
	"errors"
	"go-api-todolist/models"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoDocument = errors.New("no document with given id found")
	ErrInvalidId  = errors.New("given id is invalid")
)

func (m *mongodb) Create(ctx context.Context, t models.Todo) (string, error) {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	result, err := m.collection.InsertOne(ctx, t)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *mongodb) ReadOne(ctx context.Context, id string) (*models.Todo, error) {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidId
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	var result models.Todo
	if err := m.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNoDocument
		}
		return nil, err
	}

	return &result, nil
}

func (m *mongodb) ReadAll(ctx context.Context) ([]*models.Todo, error) {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()

	var results []*models.Todo

	cur, err := m.collection.Find(ctx, bson.D{{}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNoDocument
		}
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var elem models.Todo
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *mongodb) Update(ctx context.Context, id string, t models.Todo) error {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidId
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	updates := getTodoUpdates(t)
	update := bson.D{
		{Key: "$set", Value: updates},
	}
	_, err = m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNoDocument
		}
		return err
	}

	return nil
}

func (m *mongodb) Ping(ctx context.Context) error {
	return m.collection.Database().Client().Ping(ctx, nil)
}

func (m *mongodb) Delete(ctx context.Context, id string) error {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidId
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	_, err = m.collection.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNoDocument
		}

		return err
	}
	return nil

}

func getTodoUpdates(t any) []bson.E {
	var updates []bson.E

	v := reflect.ValueOf(t)

	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		switch val.Interface().(type) {
		case string:
			if !val.IsZero() {
				updates = append(updates, bson.E{Key: v.Type().Field(i).Tag.Get("bson"), Value: val.String()})
			}
		case uint, int:
			if !val.IsZero() {
				updates = append(updates, bson.E{Key: v.Type().Field(i).Tag.Get("bson"), Value: val.Int()})
			}
		case bool:
			updates = append(updates, bson.E{Key: v.Type().Field(i).Tag.Get("bson"), Value: val.Bool()})
		}
	}
	return updates
}
