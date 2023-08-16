package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawRepository interface {
	Insert(m bson.M) error
}

type rawRepository struct {
	coll *mongo.Collection
}

func NewRawRepository(tsClient *mongo.Client) RawRepository {
	coll := tsClient.Database("bronco").Collection("github")
	return &rawRepository{
		coll: coll,
	}
}

func (rr *rawRepository) Insert(m bson.M) error {
	_, err := rr.coll.InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}
	return err
}
