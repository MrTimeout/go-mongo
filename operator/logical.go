package operator

import "go.mongodb.org/mongo-driver/bson"

// Or is used as a logical operator to join queries and if any of them matches, then the document comes out.
func Or(d ...bson.D) bson.D {
	bsonA := make(bson.A, len(d))

	for i := range d {
		bsonA[i] = d[i]
	}

	return bson.D{{Key: "$or", Value: bsonA}}
}

// Not is used to negate the expression passed as a parameter.
func Not(d bson.D) bson.D {
	return bson.D{{Key: "$not", Value: d}}
}

// Nor is used to select all the documents that doesn't match all the query conditions.
func Nor(d ...bson.D) bson.D {
	bsonA := make(bson.A, len(d))

	for i := range d {
		bsonA[i] = d[i]
	}

	return bson.D{{Key: "$nor", Value: bsonA}}
}

// And is used as a logical operator to join queries and if all are matched, then we the document comes out.
func And(d ...bson.D) bson.D {
	bsonA := make(bson.A, len(d))

	for i := range d {
		bsonA[i] = d[i]
	}

	return bson.D{{Key: "$and", Value: bsonA}}
}
