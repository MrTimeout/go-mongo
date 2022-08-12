package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertManyResultFunc func(context.Context, *mongo.Client) (*mongo.InsertManyResult, error)

type FindFunc func(context.Context, *mongo.Client) (*any, error)

type FindOneFunc func(context.Context, *mongo.Client) (*mongo.SingleResult, error)

type DeleteByObjectIDs func(context.Context, *mongo.Client) (deleted *int64, err error)

func DoInsert[T any](db, col string, arr []T) InsertManyResultFunc {
	return func(ctx context.Context, c *mongo.Client) (*mongo.InsertManyResult, error) {
		input := make([]interface{}, len(arr))
		for i := 0; i < len(arr); i++ {
			input[i] = arr[i]
		}

		return c.Database(db).Collection(col).InsertMany(ctx, input)
	}
}

func DoInsertOne[T any](db, col string, input T) InsertManyResultFunc {
	return DoInsert(db, col, []T{input})
}

func DoFind[T any](db, col string, filter any, result *T, opts ...*options.FindOptions) FindFunc {
	return func(ctx context.Context, c *mongo.Client) (*any, error) {
		cursor, err := c.Database(db).Collection(col).Find(ctx, filter, opts...)
		if err != nil {
			return nil, err
		}

		err = cursor.All(ctx, result)
		return nil, err
	}
}

func DoFindAndUpdate[T any](db, col string, filter any, toUpdate T) FindOneFunc {
	return func(ctx context.Context, c *mongo.Client) (*mongo.SingleResult, error) {
		sr := c.Database(db).Collection(col).FindOneAndUpdate(ctx, filter, toUpdate)
		return sr, sr.Err()
	}
}

func DoDeleteByStringObjectID(db, col string, objectIDs ...string) DeleteByObjectIDs {
	return func(ctx context.Context, c *mongo.Client) (*int64, error) {
		var (
			objectIDArr = make([]primitive.ObjectID, len(objectIDs))
			err         error
		)

		for i := range objectIDs {
			objectIDArr[i], err = primitive.ObjectIDFromHex(objectIDs[i])
			if err != nil {
				return nil, err
			}
		}

		return DoDeleteByObjectID(db, col, objectIDArr...)(ctx, c)
	}
}

func DoDeleteByObjectID[T any](db, col string, objectIDs ...T) DeleteByObjectIDs {
	return func(ctx context.Context, c *mongo.Client) (*int64, error) {
		dr, err := c.Database(db).Collection(col).DeleteMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objectIDs}}}})
		if dr == nil {
			return nil, err
		}

		return &dr.DeletedCount, err
	}
}
