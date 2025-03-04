package main

import (
	"context"
	pb "github.com/Donng/shipper/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repo *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repo.collection.InsertOne(context.Background(), consignment)
	return err
}

func (repo *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repo.collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}

	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, nil
}
