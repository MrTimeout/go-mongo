package mongodb

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client     *mongo.Client
	clientErr  error
	clientOnce sync.Once
)

func GetMongoClient() (*mongo.Client, error) {
	clientOnce.Do(func() {
		ctx, cl := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cl()

		opt := options.Client().
			ApplyURI("mongodb://localhost:27017").
			SetAuth(options.Credential{Username: "admin", Password: "abc123."})

		cli, err := mongo.Connect(ctx, opt)
		if err != nil {
			clientErr = err
			return
		}

		if err := cli.Ping(ctx, readpref.Primary()); err != nil {
			clientErr = err
			return
		}

		client = cli
	})

	return client, clientErr
}

func DialConnection[T any](ctx context.Context, query func(context.Context, *mongo.Client) (*T, error)) (*T, error) {
	cli, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	result, err := query(ctx, cli)
	if err != nil {
		return nil, err
	}

	return result, nil
}
