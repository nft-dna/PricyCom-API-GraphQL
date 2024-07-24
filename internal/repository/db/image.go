package db

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MM
// ---------------------------------------------------------------

const (
	// CoImages is the name of database collection.
	coImages = "images"

	fiImageType = "type"

	fiImageContent = "content"
)

func (db *SharedMongoDbBridge) GetImage(hash string) (image *types.Image, err error) {
	col := db.client.Database(db.dbName).Collection(coImages)
	hexid, _ := types.ImageIDFromHexString(hash)
	result := col.FindOne(context.Background(), bson.D{{Key: fieldId, Value: hexid}})

	var row types.Image
	if err = result.Decode(&row); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Errorf("can not decode image %s; %s", hash, err.Error())
		return nil, err
	}
	return &row, err
}

func (db *SharedMongoDbBridge) StoreImage(image *types.Image) error {
	if image == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(coImages)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	// try to do the insert
	id := image.ID()
	rs, err := col.UpdateOne(
		ctx,
		bson.D{{Key: fieldId, Value: id}},
		bson.D{
			{Key: "$set", Value: image},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fieldId, Value: id},
			}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Errorf("can not store image %s; %s", image.ID().Hex(), err.Error())
		return err
	}
	if rs.UpsertedCount > 0 {
		log.Infof("image %s added to database", image.ID().Hex())
	}
	return nil
}

// ---------------------------------------------------------------
