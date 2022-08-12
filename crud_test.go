package mongodb

import (
	"context"
	"regexp"
	"testing"
	"time"

	o "github.com/MrTimeout/go-mongo/operator"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	crudTestDb         = "testing"
	crudTestCollection = "crud_test"
)

type Person struct {
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Age         int                `bson:"age"`
	Salary      float32            `bson:"salary"`
	SalaryDot0  float64            `bson:"salary_dot_0"`
	FavNumbers  []int              `bson:"fav_numbers"`
	BestDayEver primitive.DateTime `bson:"best_day_ever"`
}

type ComplexStruct struct {
	StringValue   string  `bson:"string_value"`
	ArrayIntValue []int   `bson:"array_int_value"`
	IntValue      int     `bson:"int_value"`
	FloatValue    float32 `bson:"float_value"`
}

func insertDocuments[T any](t *testing.T, db, col string, data []T) {
	ctx, cl := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cl()

	insertManyResult, err := DialConnection(ctx, DoInsert(db, col, data))

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		ctx, cl := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cl()
		DialConnection(ctx, DoDeleteByObjectID(crudTestDb, crudTestCollection, insertManyResult.InsertedIDs...))
	})
}

type Something int

func TestFindByComparisonOperator(t *testing.T) {
	var input = []Person{
		{
			Name:        "John",
			Surname:     "Sullivan",
			Age:         5,
			Salary:      20.0,
			FavNumbers:  []int{1, 2, 3},
			BestDayEver: primitive.NewDateTimeFromTime(time.Now().Add(-24 * time.Hour)),
		},
		{
			Name:        "Ivan",
			Surname:     "Martinez Alberte",
			Age:         24,
			Salary:      1500.0,
			FavNumbers:  []int{23, 73},
			BestDayEver: primitive.NewDateTimeFromTime(time.Now().Add(-48 * time.Hour)),
		},
		{
			Name:        "Pedro",
			Surname:     "Gonzalez Gonzalez",
			Age:         66,
			Salary:      2000.0,
			FavNumbers:  []int{101, 200},
			BestDayEver: primitive.NewDateTimeFromTime(time.Now().Add(-72 * time.Hour)),
		},
	}

	var testCases = []struct {
		description string
		filter      bson.D
		want        []Person
	}{
		{
			description: "gt only returns the greater than the number passed as a parameter",
			filter:      o.F("age", o.Gt(24)), // bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 24}}}}
			want:        input[2:],
		},
		{
			description: "gt only returns the documents which has an element, which is inside the array, greater than the value passed as a parameter",
			filter:      o.F("fav_numbers", o.Gt(50)),
			want:        input[1:],
		},
		{
			description: "gt only returns the documents which has a date greater than the date passed as a parameter",
			filter:      o.F("best_day_ever", o.Gt(time.Now().Add(-36*time.Hour))),
			want:        input[:1],
		},
		{
			description: "gt only returns the documents which has a float greater than the actual passed as a parameter",
			filter:      o.F("salary", o.Gt(1400.0)),
			want:        input[1:],
		},
		{
			description: "gte only returns the greater than or equal the number passed as a parameter",
			filter:      o.F("age", o.Gte(24)), // bson.D{{Key: "age", Value: bson.D{{Key: "$gte", Value: 24}}}}
			want:        input[1:],
		},
		{
			description: "lt only returns the lower than the number passed as a parameter",
			filter:      o.F("age", o.Lt(24)), // bson.D{{Key: "age", Value: bson.D{{Key: "$lt", Value: 24}}}}
			want:        input[:1],
		},
		{
			description: "lte only returns the lower than or equal the number passed as a parameter",
			filter:      o.F("age", o.Lte(24)), // bson.D{{Key: "age", Value: bson.D{{Key: "$lte", Value: 24}}}}
			want:        input[:2],
		},
		{
			description: "gt and lt only returns the lower than and greater than the number passed as a parameter",
			filter:      o.F("age", o.Gt(23), o.Lt(70)), // bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 24}, {Key: "$lt", Value:24}}}}
			want:        input[1:],
		},
		{
			description: "gte and lte only returns the lower than or equal and greater than or equal the number passed as a parameter",
			filter:      o.F("age", o.Gte(5), o.Lte(66)), // bson.D{{Key: "age", Value: bson.D{{Key: "$gte", Value: 24}, {Key: "$lte", Value: 24}}}}
			want:        input,
		},
		{
			description: "gt and lte only returns the lower than or equal and greater than the number passed as a parameter",
			filter:      o.F("age", o.Gt(24), o.Lte(66)), // bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 24}, {Key: "$lte", Value: 24}}}}
			want:        input[2:],
		},
		{
			description: "gte and lt only returns the lower than and greater than or equal the number passed as a parameter",
			filter:      o.F("age", o.Gte(24), o.Lt(66)), // bson.D{{Key: "age", Value: bson.D{{Key: "$gte", Value: 24}, {Key: "$lt", Value: 24}}}}
			want:        []Person{input[1]},
		},
		{
			description: "eq returns the elements that are equal to the value passed as a parameter when int",
			filter:      o.F("age", o.Eq(24)), // bson.D{{Key: "age", Value: bson.D{{Key: "$eq", Value: 24}}}}
			want:        []Person{input[1]},
		},
		{
			description: "eq returns the elements that are equal to the value passed as a parameter when string",
			filter:      o.F("name", o.Eq("Ivan")), // bson.D{{Key: "name", Value: bson.D{{Key: "$eq", Value: "Ivan"}}}}
			want:        []Person{input[1]},
		},
		{
			description: "eq returns the elements that are equal to the value passed as a parameter when float32",
			filter:      o.F("salary", o.Eq(2000.0)), // bson.D{{Key: "salary", Value: bson.D{{Key: "$eq", Value: 2000.0}}}}
			want:        input[2:],
		},
		{
			description: "ne retuns the elements that are not equal to the value passed as a parameter when int",
			filter:      o.F("age", o.Ne(3)), // bson.D{{Key: "age", Value: bson.D{{Key: "$ne", Value: 3}}}}
			want:        input,
		},
		{
			description: "ne retuns the elements that are not equal to the value passed as a parameter when string",
			filter:      o.F("name", o.Ne("Pedro")), // bson.D{{Key: "name", Value: bson.D{{Key: "$ne", Value: "Pedro"}}}}
			want:        input[:2],
		},
		{
			description: "ne retuns the elements that are not equal to the value passed as a parameter when float32",
			filter:      o.F("salary", o.Ne(2.0)), // bson.D{{Key: "salary", Value: bson.D{{Key: "$ne", Value: 2.0}}}}
			want:        input,
		},
	}

	ctx, cl := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cl()

	insertDocuments(t, crudTestDb, crudTestCollection, input)

	for _, tCase := range testCases {
		t.Run(tCase.description, func(t *testing.T) {
			got := make([]Person, 3)

			_, err := DialConnection(ctx, DoFind(crudTestDb, crudTestCollection, tCase.filter, &got))
			if err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, tCase.want, got)
		})
	}
}

func TestComplexFindQueries(t *testing.T) {
	var input = []ComplexStruct{
		{
			StringValue:   "Hello world",
			IntValue:      10,
			ArrayIntValue: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			FloatValue:    2134.4,
		},
		{
			StringValue:   "Hello world",
			IntValue:      30,
			ArrayIntValue: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			FloatValue:    2134.4,
		},
		{
			StringValue:   "Hello world",
			IntValue:      20,
			ArrayIntValue: []int{1, 2},
			FloatValue:    1000.4,
		},
		{
			StringValue:   "Bye world",
			IntValue:      20,
			ArrayIntValue: []int{1, 2},
			FloatValue:    1000.4,
		},
		{
			StringValue:   "Hello world",
			IntValue:      40,
			ArrayIntValue: []int{1, 2, 3, 4},
			FloatValue:    2134.4,
		},
		{
			StringValue:   "Hello world",
			IntValue:      40,
			ArrayIntValue: []int{1, 2, 3},
			FloatValue:    2134.4,
		},
	}
	var testCases = []struct {
		description string
		filter      bson.D
		want        []ComplexStruct
	}{
		{
			description: "filter by string, array, float and int value using comparison operators being the int value the difference",
			filter: o.And(
				o.F("int_value", o.Gte(5), o.In([]int{1, 2, 3, 10})),
				o.F("float_value", o.Gt(100)),
				o.F("array_int_value", o.Eq(1)),
				o.F("string_value", o.Eq("Hello world")),
			),
			want: []ComplexStruct{input[0]},
		},
		{
			description: "filter by string, array, float and int value using comparison operator being the string value the difference",
			filter: o.And(
				o.F("int_value", o.Lt(100), o.Gt(5), o.In([]int{1, 10, 20, 30})),
				o.F("string_value", o.Regex(regexp.MustCompile("^Bye.*$"))),
				o.F("array_int_value", o.Eq(2)),
				o.F("float_value", o.Lte(2000), o.Gt(1000)),
			),
			want: []ComplexStruct{input[3]},
		},
		{
			description: "filter by array operator in",
			filter: o.And(
				o.F("int_value", o.Lte(200), o.Nin([]int{10, 50, 70})),
				o.F("string_value", o.Regex(regexp.MustCompile("^Bye.*$"))),
				o.F("float_value", o.Gt(500)),
			),
			want: []ComplexStruct{input[3]},
		},
	}

	ctx, cl := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cl()

	insertDocuments(t, crudTestDb, crudTestCollection, input)

	for _, tCase := range testCases {
		t.Run(tCase.description, func(t *testing.T) {
			result := make([]ComplexStruct, 1)

			_, err := DialConnection(ctx, DoFind(crudTestDb, crudTestCollection, tCase.filter, &result))
			if err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, tCase.want, result)
		})
	}
}
