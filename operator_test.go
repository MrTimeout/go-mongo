package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

//TODO we have to fix this test
func TestEq(t *testing.T) {
	var testCases = []struct {
		description, field string
		value              any
		want               bson.E
	}{
		{
			description: "comparison using int",
			field:       "int_value",
			value:       2,
			want:        bson.E{Key: "int_value", Value: bson.D{{Key: "$eq", Value: 2}}},
		},
		{
			description: "comparison using float",
			field:       "float_value",
			value:       2.4,
			want:        bson.E{Key: "float_value", Value: bson.D{{Key: "$eq", Value: 2.4}}},
		},
		{
			description: "comparison using string",
			field:       "string_value",
			value:       "this is the value",
			want:        bson.E{Key: "string_value", Value: bson.D{{Key: "$eq", Value: "this is the value"}}},
		},
		{
			description: "comparison using array of ints",
			field:       "array_int_value",
			value:       []int{1, 2, 3, 4},
			want:        bson.E{Key: "array_int_value", Value: bson.D{{Key: "$eq", Value: []int{1, 2, 3, 4}}}},
		},
		{
			description: "comparison using array of floats32",
			field:       "array_float32_value",
			value:       []float32{1.5, 2.3, 3.5, 4.4},
			want:        bson.E{Key: "array_float32_value", Value: bson.D{{Key: "$eq", Value: []float32{1.5, 2.3, 3.5, 4.4}}}},
		},
		{
			description: "comparison using array of strings",
			field:       "array_string_value",
			value:       []string{"first", "second", "third"},
			want:        bson.E{Key: "array_string_value", Value: bson.D{{Key: "$eq", Value: []string{"first", "second", "third"}}}},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.description, func(t *testing.T) {
			assert.Equal(t, tCase.want, eq(tCase.value))
		})
	}
}
