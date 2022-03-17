package mongo

import (
	"context"
	"go-practice/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (domain.Repository, error) {
	repo := &mongoRepository{
		timeout: time.Duration(mongoTimeout) * time.Second,
		db:      mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (mr *mongoRepository) Find(name string) (*domain.Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	metadata := &domain.Metadata{}
	collection := mr.client.Database(mr.db).Collection("metadata")
	filter := bson.M{"name": name}
	err := collection.FindOne(ctx, filter).Decode(&metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func (mr *mongoRepository) Store() error {
	return nil
}
