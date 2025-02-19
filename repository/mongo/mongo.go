package mongo

import (
	"context"
	"go-api-todolist/config"
	"go-api-todolist/repository"
	"log"
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

func New(ctx context.Context, cfg *config.Database) (repository.MongoDB, error) {
	var repo repository.MongoDB
	var err error
	connect := func() (repository.MongoDB, error) {
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

	// Doing a 3 time retry for connecting to database, each time with 5,10,15 sec wait
	for i := 0; err != nil || i == 0; i++ {
		log.Printf("mongodb connect try %d", i)
		repo, err = connect()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5 * time.Duration(i))
		}
		if i == 3 {
			log.Println("Failed to connect to database")
			return nil, err
		}
	}

	return repo, err
}

func uri(cfg *config.Database) string {
	return strings.Replace(cfg.Url, "<password>", url.QueryEscape(cfg.Password), 1)
}

func NewMongoContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Second*10)
}
