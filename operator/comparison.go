// comparison file contains all the comparison operators which are used in mongo to compare values.
//
// $gt, $gte, $lt, $lte can be used with:
//	- int8, uint8, int16, uin16, int32, uint32, int64, uint64, float32, float64, time.Time, primitive.ObjectId
//	- []int8, []uin8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []float32, []float64, []time.Time
package operator

import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ComparableThanType is the type of values which can be passed to some comparison functions like Gt, Gte, Lt, Lte
type ComparableThanType interface {
	int8 | uint8 | int16 | uint16 | int | int32 | uint32 | int64 | uint64 | float32 | float64 | time.Time |
		[]int8 | []uint8 | []int16 | []uint16 | []int | []int32 | []uint32 | []int64 | []uint64 | []float32 | []float64 | []time.Time
}

func F(field string, conds ...bson.E) bson.D {
	bsonD := make(bson.D, len(conds))

	for i := range conds {
		bsonD[i] = conds[i]
	}

	return bson.D{{Key: field, Value: bsonD}}
}

// Eq returns a { "$eq": "<value>" }
func Eq(value any) bson.E {
	switch p := value.(type) {
	case time.Time:
		return bson.E{Key: "$ne", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$eq", Value: value}
}

// Ne returns a { "$ne": "<value>" }
func Ne(value any) bson.E {
	switch p := value.(type) {
	case time.Time:
		return bson.E{Key: "$ne", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$ne", Value: value}
}

// Gt returns a { "$gt": "<value>" }
func Gt[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$gt", Value: primitive.NewDateTimeFromTime(p)}
	case []uint8:
		// We match primitive.ObjectId
		if len(p) == 12 {
			objectId, err := primitive.ObjectIDFromHex(string(p))
			if err != nil {
				return bson.E{Key: "$gt", Value: objectId}
			}
		}
	}
	return bson.E{Key: "$gt", Value: value}
}

// Gte returns a { "$gte": "<value>" }
func Gte[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$gte", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$gte", Value: value}
}

// Lt returns a { "$lt": "<value>" }
func Lt[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$lt", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$lt", Value: value}
}

// Lte returns a { "$lte": "<value>" }
func Lte[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$lte", Value: value}
}

// Regex returns a { "$regex": /<value>/ }
func Regex[T *regexp.Regexp | string](value T) bson.E {
	var val string
	switch p := any(value).(type) {
	case *regexp.Regexp:
		val = p.String()
	case string:
		val = p
	}
	return bson.E{Key: "$regex", Value: val}
}

// In returns a { "$in": Array<any>[] }
func In[T any](value []T) bson.E {
	var arr = make([]any, len(value))

	doInclusionOperator(value, arr)

	return bson.E{Key: "$in", Value: value}
}

// Nin returns a { "$nin": Array<any>[] }
func Nin[T any](value []T) bson.E {
	var arr = make([]any, len(value))

	doInclusionOperator(value, arr)

	return bson.E{Key: "$nin", Value: value}
}

// TODO we have to create a function that returns this encoders
// so we can loop accross it using them, without switching each
// iteration.
func doInclusionOperator[T any](value []T, arr []any) {
	for i := range value {
		switch p := any(value[i]).(type) {
		case time.Time:
			arr[i] = primitive.NewDateTimeFromTime(p)
		case *regexp.Regexp:
			arr[i] = p.String()
		default:
			arr[i] = p
		}
	}
}

// Type returns a { "$type": "<value>" }
func Type(value any) bson.E {
	return bson.E{Key: "$type", Value: value}
}
