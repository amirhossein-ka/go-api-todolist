package mongo

import (
	"context"
	"go-api-todolist/config"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongodb struct {
	collection *mongo.Collection
}

func New(ctx context.Context, cfg *config.Database) (*mongodb, error) {
	ctx, cancel := NewMongoContext(ctx)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri(cfg)))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	col := client.Database(cfg.DB).Collection(cfg.Collection)

	return &mongodb{collection: col}, nil
}

func uri(cfg *config.Database) string {
	return strings.Replace(cfg.Url, "<password>", url.QueryEscape(cfg.Password), 1)
}

func NewMongoContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Second*10)
}
