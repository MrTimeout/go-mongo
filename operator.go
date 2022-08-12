package mongodb

import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func f(field string, conds ...bson.E) bson.D {
	bsonD := make(bson.D, len(conds))

	for i := range conds {
		bsonD[i] = conds[i]
	}

	return bson.D{{Key: field, Value: bsonD}}
}

//eq https://www.mongodb.com/docs/manual/reference/operator/query/eq/#mongodb-query-op.-eq
func eq(value any) bson.E {
	return bson.E{Key: "$eq", Value: value}
}

func ne(value any) bson.E {
	return bson.E{Key: "$ne", Value: value}
}

func gt[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$gt", Value: value}
}

func gte[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$gte", Value: value}
}

func lt[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$lt", Value: value}
}

func lte[T ComparableThanType](value T) bson.E {
	switch p := any(value).(type) {
	case time.Time:
		return bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(p)}
	}
	return bson.E{Key: "$lte", Value: value}
}

func regex[T *regexp.Regexp | string](value T) bson.E {
	var val string
	switch p := any(value).(type) {
	case *regexp.Regexp:
		val = p.String()
	case string:
		val = p
	}
	return bson.E{Key: "$regex", Value: val}
}

func in[T any](value []T) bson.E {
	return bson.E{Key: "$in", Value: value}
}

func nin[T any](value []T) bson.E {
	return bson.E{Key: "$nin", Value: value}
}

func isType(value any) bson.E {
	return bson.E{Key: "$type", Value: value}
}

func or(d ...bson.D) bson.D {
	bsonA := make(bson.A, len(d))

	for i := range d {
		bsonA[i] = d[i]
	}

	return bson.D{{Key: "$or", Value: bsonA}}
}

func and(d ...bson.D) bson.D {
	bsonA := make(bson.A, len(d))

	for i := range d {
		bsonA[i] = d[i]
	}

	return bson.D{{Key: "$and", Value: bsonA}}
}

type ComparableThanType interface {
	int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | time.Time
}
