package mongo

import (
	"context"
	"fmt"
	"go-api-todolist/models"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongodb) Create(ctx context.Context, t models.Todo) (any, error) {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	result, err := m.collection.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}

	fmt.Println(result.InsertedID)
	return result, nil
}

func (m *mongodb) ReadOne(ctx context.Context, id uint) (*models.Todo, error) {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()

	filter := bson.D{primitive.E{Key: "name", Value: id}}
	var result models.Todo
	if err := m.collection.FindOne(ctx, filter).Decode(&result); err != nil {
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

func (m *mongodb) Update(ctx context.Context, id uint, t *models.Todo) error {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	filter := bson.D{primitive.E{Key: "id", Value: id}}

	updates := getTodoUpdates(t)
	update := bson.D{
		{Key: "$set", Value: updates},
	}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongodb) Delete(ctx context.Context, id uint) error {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()

	filter := bson.D{primitive.E{Key: "id", Value: id}}

	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil

}

func getTodoUpdates(t any) []bson.E {
	var updates []bson.E

	v := reflect.ValueOf(t).Elem()
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
