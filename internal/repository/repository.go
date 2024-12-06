package repository

import (
	"context"
	"product/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetAll()([]models.Product, error)
	Create(product models.Product) error
	Delete(id string) error
	Update(id string , updateditem models.Product)(models.Product, error)
}


type MongoRepository struct{
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, dbName, collection string) Repository{
	return &MongoRepository{
		collection: client.Database(dbName).Collection(collection),
	}
}

func (r *MongoRepository)GetAll()([]models.Product, error){
	ctx , cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	cursor , err := r.collection.Find(ctx,bson.D{})
	if err!= nil{
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx,&products); err!=nil{
		return nil, err
	}
	return products, nil

}

func (r *MongoRepository)Create(product models.Product) error{
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	_,err := r.collection.InsertOne(ctx, &product)
	return err
}

func (r *MongoRepository)Delete(id string) error{
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	objid , _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objid})
	return err
}

func (r *MongoRepository)Update(id string , updatedproduct models.Product)(models.Product, error){
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	objid, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set": updatedproduct,
	}
	result, err := r.collection.UpdateOne(ctx,bson.M{"_id": objid},update)
	if err != nil{
		return models.Product{}, err
	}
	if result.ModifiedCount ==0 {
		return models.Product{}, mongo.ErrNoDocuments
	}
	return updatedproduct, nil
}
