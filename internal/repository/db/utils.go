// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// exists checks if the given filter matches at least one document in the given DB collection.
func (db *SharedMongoDbBridge) exists(col *mongo.Collection, filter *bson.D) bool {
	sr := col.FindOne(
		context.Background(),
		filter,
		options.FindOne().SetProjection(bson.D{{Key: fieldId, Value: true}}),
	)
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("check failed; %s", sr.Err().Error())
		}
		return false
	}
	return true
}

// exists checks if the given filter matches at least one document in the given DB collection.
func (db *MongoDbBridge) exists(col *mongo.Collection, filter *bson.D) bool {
	sr := col.FindOne(
		context.Background(),
		filter,
		options.FindOne().SetProjection(bson.D{{Key: fieldId, Value: true}}),
	)
	if sr.Err() != nil {
		if sr.Err() != mongo.ErrNoDocuments {
			log.Errorf("check failed; %s", sr.Err().Error())
		}
		return false
	}
	return true
}

// AggregateSingle reads and parses single row aggregation pipeline into provided value structure.
func (db *MongoDbBridge) AggregateSingle(col *mongo.Collection, pipe *mongo.Pipeline, val interface{}) error {
	// get the lowest listing time
	agg, err := col.Aggregate(context.Background(), *pipe)
	if err != nil {
		return err
	}

	defer closeFindCursor("AggregateSingle", agg)

	if !agg.Next(context.Background()) {
		return mongo.ErrNoDocuments
	}

	if err := agg.Decode(val); err != nil {
		return err
	}
	return nil
}

// filterAddDateTimeLimit adds a time limit of type <wall> into the provided bson.D filter.
// I.e.: <field> exists AND is not null AND <operand> <time stamp>
func filterAddDateTimeLimit(filter bson.D, field string, operand string, ts types.Time) bson.D {
	return append(filter, bson.E{Key: field, Value: bson.D{
		{Key: "$exists", Value: true},
		{Key: "$ne", Value: nil},
		{Key: operand, Value: ts},
	}})
}

// filterAddDateTimeMiss adds a missed time limit of type <wall> into the provided bson.D filter.
// I.e.: <field> not exists OR is null OR <operand> <time stamp>
func filterAddDateTimeMiss(filter bson.D, field string, operand string, ts types.Time) bson.D {
	return append(filter, bson.E{Key: "$or", Value: bson.A{
		bson.D{{Key: field, Value: bson.D{{Key: "$exists", Value: false}}}},
		bson.D{{Key: field, Value: bson.D{{Key: "$eq", Value: nil}}}},
		bson.D{{Key: field, Value: bson.D{{Key: operand, Value: ts}}}},
	}})
}
