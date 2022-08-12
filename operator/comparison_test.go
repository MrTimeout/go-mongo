package operator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestEqAndNe(t *testing.T) {
	var testCases = []struct {
		description, field string
		value              any
		wantEq             bson.D
		wantNe             bson.D
	}{
		{
			description: "comparison using int32",
			field:       "int32_value",
			value:       2,
			wantEq:      bson.D{{Key: "int32_value", Value: bson.D{{Key: "$eq", Value: 2}}}},
			wantNe:      bson.D{{Key: "int32_value", Value: bson.D{{Key: "$ne", Value: 2}}}},
		},
		{
			description: "comparison using int64",
			field:       "int64_value",
			value:       int64(4),
			wantEq:      bson.D{{Key: "int64_value", Value: bson.D{{Key: "$eq", Value: int64(4)}}}},
			wantNe:      bson.D{{Key: "int64_value", Value: bson.D{{Key: "$ne", Value: int64(4)}}}},
		},
		{
			description: "comparison using float",
			field:       "float_value",
			value:       2.4,
			wantEq:      bson.D{{Key: "float_value", Value: bson.D{{Key: "$eq", Value: 2.4}}}},
			wantNe:      bson.D{{Key: "float_value", Value: bson.D{{Key: "$ne", Value: 2.4}}}},
		},
		{
			description: "comparison using bool",
			field:       "bool_value",
			value:       true,
			wantEq:      bson.D{{Key: "bool_value", Value: bson.D{{Key: "$eq", Value: true}}}},
			wantNe:      bson.D{{Key: "bool_value", Value: bson.D{{Key: "$ne", Value: true}}}},
		},
		{
			description: "comparison using string",
			field:       "string_value",
			value:       "this is the value",
			wantEq:      bson.D{{Key: "string_value", Value: bson.D{{Key: "$eq", Value: "this is the value"}}}},
			wantNe:      bson.D{{Key: "string_value", Value: bson.D{{Key: "$ne", Value: "this is the value"}}}},
		},
		{
			description: "comparison using array of int32",
			field:       "array_int32_value",
			value:       []int{1, 2, 3, 4},
			wantEq:      bson.D{{Key: "array_int32_value", Value: bson.D{{Key: "$eq", Value: []int{1, 2, 3, 4}}}}},
			wantNe:      bson.D{{Key: "array_int32_value", Value: bson.D{{Key: "$ne", Value: []int{1, 2, 3, 4}}}}},
		},
		{
			description: "comparison using array of int64",
			field:       "array_int64_value",
			value:       []int64{1, 2, 3, 4},
			wantEq:      bson.D{{Key: "array_int64_value", Value: bson.D{{Key: "$eq", Value: []int64{1, 2, 3, 4}}}}},
			wantNe:      bson.D{{Key: "array_int64_value", Value: bson.D{{Key: "$ne", Value: []int64{1, 2, 3, 4}}}}},
		},
		{
			description: "comparison using array of float32",
			field:       "array_float32_value",
			value:       []float32{1.5, 2.3, 3.5, 4.4},
			wantEq:      bson.D{{Key: "array_float32_value", Value: bson.D{{Key: "$eq", Value: []float32{1.5, 2.3, 3.5, 4.4}}}}},
			wantNe:      bson.D{{Key: "array_float32_value", Value: bson.D{{Key: "$ne", Value: []float32{1.5, 2.3, 3.5, 4.4}}}}},
		},
		{
			description: "comparison using array of bool",
			field:       "array_bool_value",
			value:       []bool{true, false},
			wantEq:      bson.D{{Key: "array_bool_value", Value: bson.D{{Key: "$eq", Value: []bool{true, false}}}}},
			wantNe:      bson.D{{Key: "array_bool_value", Value: bson.D{{Key: "$ne", Value: []bool{true, false}}}}},
		},
		{
			description: "comparison using array of strings",
			field:       "array_string_value",
			value:       []string{"first", "second", "third"},
			wantEq:      bson.D{{Key: "array_string_value", Value: bson.D{{Key: "$eq", Value: []string{"first", "second", "third"}}}}},
			wantNe:      bson.D{{Key: "array_string_value", Value: bson.D{{Key: "$ne", Value: []string{"first", "second", "third"}}}}},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.description, func(t *testing.T) {
			assert.Equal(t, tCase.wantEq, F(tCase.field, Eq(tCase.value)))
			assert.Equal(t, tCase.wantNe, F(tCase.field, Ne(tCase.value)))
		})
	}
}

func TestComparisonGtGteLtLte(t *testing.T) {
	var testCases = []struct {
		description, field string
		value              any
		wantGt             bson.D
		wantGte            bson.D
		wantLt             bson.D
		wantLte            bson.D
	}{
		{
			description: "comparison using ObjectId as array of uint8/byte",
			field:       "objectId_value",
			value:       []byte("572bb8222b288919b68abf6b"),
			wantGt:      bson.D{{Key: "objectId_value", Value: bson.D{{Key: "$gt", Value: objectIdHelper(t, []byte("572bb8222b288919b68abf6b"))}}}},
			wantGte:     bson.D{{Key: "objectId_value", Value: bson.D{{Key: "$gte", Value: objectIdHelper(t, []byte("572bb8222b288919b68abf6b"))}}}},
			wantLt:      bson.D{{Key: "objectId_value", Value: bson.D{{Key: "$lt", Value: objectIdHelper(t, []byte("572bb8222b288919b68abf6b"))}}}},
			wantLte:     bson.D{{Key: "objectId_value", Value: bson.D{{Key: "$lte", Value: objectIdHelper(t, []byte("572bb8222b288919b68abf6b"))}}}},
		},
		{
			description: "comparison using int8",
			field:       "int8_value",
			value:       int8(2),
			wantGt:      bson.D{{Key: "int8_value", Value: bson.D{{Key: "$gt", Value: int8(2)}}}},
			wantGte:     bson.D{{Key: "int8_value", Value: bson.D{{Key: "$gte", Value: int8(2)}}}},
			wantLt:      bson.D{{Key: "int8_value", Value: bson.D{{Key: "$lt", Value: int8(2)}}}},
			wantLte:     bson.D{{Key: "int8_value", Value: bson.D{{Key: "$lte", Value: int8(2)}}}},
		},
		{
			description: "comparison using uint8",
			field:       "uint8_value",
			value:       uint8(2),
			wantGt:      bson.D{{Key: "uint8_value", Value: bson.D{{Key: "$gt", Value: uint8(2)}}}},
			wantGte:     bson.D{{Key: "uint8_value", Value: bson.D{{Key: "$gte", Value: uint8(2)}}}},
			wantLt:      bson.D{{Key: "uint8_value", Value: bson.D{{Key: "$lt", Value: uint8(2)}}}},
			wantLte:     bson.D{{Key: "uint8_value", Value: bson.D{{Key: "$lte", Value: uint8(2)}}}},
		},
		{
			description: "comparison using int16",
			field:       "int16_value",
			value:       int16(2),
			wantGt:      bson.D{{Key: "int16_value", Value: bson.D{{Key: "$gt", Value: int16(2)}}}},
			wantGte:     bson.D{{Key: "int16_value", Value: bson.D{{Key: "$gte", Value: int16(2)}}}},
			wantLt:      bson.D{{Key: "int16_value", Value: bson.D{{Key: "$lt", Value: int16(2)}}}},
			wantLte:     bson.D{{Key: "int16_value", Value: bson.D{{Key: "$lte", Value: int16(2)}}}},
		},
		{
			description: "comparison using uint16",
			field:       "uint16_value",
			value:       uint16(2),
			wantGt:      bson.D{{Key: "uint16_value", Value: bson.D{{Key: "$gt", Value: uint16(2)}}}},
			wantGte:     bson.D{{Key: "uint16_value", Value: bson.D{{Key: "$gte", Value: uint16(2)}}}},
			wantLt:      bson.D{{Key: "uint16_value", Value: bson.D{{Key: "$lt", Value: uint16(2)}}}},
			wantLte:     bson.D{{Key: "uint16_value", Value: bson.D{{Key: "$lte", Value: uint16(2)}}}},
		},
		{
			description: "comparison using int32",
			field:       "int32_value",
			value:       2,
			wantGt:      bson.D{{Key: "int32_value", Value: bson.D{{Key: "$gt", Value: 2}}}},
			wantGte:     bson.D{{Key: "int32_value", Value: bson.D{{Key: "$gte", Value: 2}}}},
			wantLt:      bson.D{{Key: "int32_value", Value: bson.D{{Key: "$lt", Value: 2}}}},
			wantLte:     bson.D{{Key: "int32_value", Value: bson.D{{Key: "$lte", Value: 2}}}},
		},
		{
			description: "comparison using uint32",
			field:       "uint32_value",
			value:       2,
			wantGt:      bson.D{{Key: "uint32_value", Value: bson.D{{Key: "$gt", Value: uint32(2)}}}},
			wantGte:     bson.D{{Key: "uint32_value", Value: bson.D{{Key: "$gte", Value: uint32(2)}}}},
			wantLt:      bson.D{{Key: "uint32_value", Value: bson.D{{Key: "$lt", Value: uint32(2)}}}},
			wantLte:     bson.D{{Key: "uint32_value", Value: bson.D{{Key: "$lte", Value: uint32(2)}}}},
		},
		{
			description: "comparison using int64",
			field:       "int64_value",
			value:       int64(4),
			wantGt:      bson.D{{Key: "int64_value", Value: bson.D{{Key: "$gt", Value: int64(4)}}}},
			wantGte:     bson.D{{Key: "int64_value", Value: bson.D{{Key: "$gte", Value: int64(4)}}}},
			wantLt:      bson.D{{Key: "int64_value", Value: bson.D{{Key: "$lt", Value: int64(4)}}}},
			wantLte:     bson.D{{Key: "int64_value", Value: bson.D{{Key: "$lte", Value: int64(4)}}}},
		},
		{
			description: "comparison using uint64",
			field:       "uint64_value",
			value:       uint64(4),
			wantGt:      bson.D{{Key: "uint64_value", Value: bson.D{{Key: "$gt", Value: uint64(4)}}}},
			wantGte:     bson.D{{Key: "uint64_value", Value: bson.D{{Key: "$gte", Value: uint64(4)}}}},
			wantLt:      bson.D{{Key: "uint64_value", Value: bson.D{{Key: "$lt", Value: uint64(4)}}}},
			wantLte:     bson.D{{Key: "uint64_value", Value: bson.D{{Key: "$lte", Value: uint64(4)}}}},
		},
		{
			description: "comparison using float",
			field:       "float_value",
			value:       2.4,
			wantGt:      bson.D{{Key: "float_value", Value: bson.D{{Key: "$gt", Value: 2.4}}}},
			wantGte:     bson.D{{Key: "float_value", Value: bson.D{{Key: "$gte", Value: 2.4}}}},
			wantLt:      bson.D{{Key: "float_value", Value: bson.D{{Key: "$lt", Value: 2.4}}}},
			wantLte:     bson.D{{Key: "float_value", Value: bson.D{{Key: "$lte", Value: 2.4}}}},
		},
		{
			description: "comparison using string",
			field:       "string_value",
			value:       "this is the value",
			wantGt:      bson.D{{Key: "string_value", Value: bson.D{{Key: "$eq", Value: "this is the value"}}}},
			wantGte:     bson.D{{Key: "string_value", Value: bson.D{{Key: "$ne", Value: "this is the value"}}}},
		},
		{
			description: "comparison using array of int32",
			field:       "array_int32_value",
			value:       []int{1, 2, 3, 4},
			wantGt:      bson.D{{Key: "array_int32_value", Value: bson.D{{Key: "$eq", Value: []int{1, 2, 3, 4}}}}},
			wantGte:     bson.D{{Key: "array_int32_value", Value: bson.D{{Key: "$ne", Value: []int{1, 2, 3, 4}}}}},
		},
		{
			description: "comparison using array of int64",
			field:       "array_int64_value",
			value:       []int64{1, 2, 3, 4},
			wantGt:      bson.D{{Key: "array_int64_value", Value: bson.D{{Key: "$eq", Value: []int64{1, 2, 3, 4}}}}},
			wantGte:     bson.D{{Key: "array_int64_value", Value: bson.D{{Key: "$ne", Value: []int64{1, 2, 3, 4}}}}},
		},
		{
			description: "comparison using array of float32",
			field:       "array_float32_value",
			value:       []float32{1.5, 2.3, 3.5, 4.4},
			wantGt:      bson.D{{Key: "array_float32_value", Value: bson.D{{Key: "$eq", Value: []float32{1.5, 2.3, 3.5, 4.4}}}}},
			wantGte:     bson.D{{Key: "array_float32_value", Value: bson.D{{Key: "$ne", Value: []float32{1.5, 2.3, 3.5, 4.4}}}}},
		},
		{
			description: "comparison using array of bool",
			field:       "array_bool_value",
			value:       []bool{true, false},
			wantGt:      bson.D{{Key: "array_bool_value", Value: bson.D{{Key: "$eq", Value: []bool{true, false}}}}},
			wantGte:     bson.D{{Key: "array_bool_value", Value: bson.D{{Key: "$ne", Value: []bool{true, false}}}}},
		},
		{
			description: "comparison using array of strings",
			field:       "array_string_value",
			value:       []string{"first", "second", "third"},
			wantGt:      bson.D{{Key: "array_string_value", Value: bson.D{{Key: "$eq", Value: []string{"first", "second", "third"}}}}},
			wantGte:     bson.D{{Key: "array_string_value", Value: bson.D{{Key: "$ne", Value: []string{"first", "second", "third"}}}}},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.description, func(t *testing.T) {
			assert.Equal(t, tCase.wantGt, F(tCase.field, Eq(tCase.value)))
			assert.Equal(t, tCase.wantGte, F(tCase.field, Ne(tCase.value)))
		})
	}
}

func objectIdHelper(t *testing.T, input []byte) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(string(input))
	if err != nil {
		t.Fatal(err)
	}

	return objectId
}
