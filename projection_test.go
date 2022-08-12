package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	db                   = "testing"
	projectionCollection = "projection_test"
)

type longDumbStruct struct {
	First  string  `bson:"first"`
	Second int64   `bson:"second"`
	Third  int     `bson:"third"`
	Fourth float32 `bson:"fourth"`
	Fifth  bool    `bson:"fifth"`
	Sixth  []int   `bson:"sixth"`
}

func TestProjection(t *testing.T) {
	t.Run("testing projection of inclusion of fields", func(t *testing.T) {
		ctx, cl := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cl()
		var result []struct {
			First string `bson:"first"`
			Third int    `bson:"third"`
		}
		filter := f("first", eq("hello"))
		options := options.Find().SetProjection(bson.D{{Key: "first", Value: 1}, {Key: "third", Value: 1}, {Key: "Fifth", Value: 1}, {Key: "_id", Value: 0}})

		insertDocuments(t, db, projectionCollection, []longDumbStruct{
			{
				First:  "hello",
				Second: 2,
				Third:  4,
				Fourth: 6.7,
				Fifth:  true,
				Sixth:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			{
				First:  "bye",
				Second: 1,
				Third:  10,
				Fourth: 20.7,
				Fifth:  false,
				Sixth:  []int{8, 7, 6, 5, 4, 3, 2, 1},
			},
		})

		_, err := DialConnection(ctx, DoFind(db, projectionCollection, filter, &result, options))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "hello", result[0].First)
		assert.Equal(t, 4, result[0].Third)
	})
}
