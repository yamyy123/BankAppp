package service

import (
	"bankDemo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Loans struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitLoan(collection *mongo.Collection, ctx context.Context) *Loans{
	return &Loans{ctx,collection}
}

func (l *Loans)CreateLoan(user *models.Loan) (*mongo.InsertOneResult, error){
	indexModel:= []mongo.IndexModel{
		{
			Keys: bson.M{"customer_id": 1}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := l.mongoCollection.Indexes().CreateMany(l.ctx, indexModel)
	if err != nil {
		return nil,err
	}
	result,err := l.mongoCollection.InsertOne(l.ctx,user)
	if(err!=nil){
		return nil,err
	}
	return result,nil
}
func (l *Loans)GetLoanById(id int64) (*models.Loan, error){
	filter := bson.M{"_id": id}
	var loan models.Loan
	err := l.mongoCollection.FindOne(l.ctx, filter).Decode(&loan)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}
func (l *Loans)UpdateLoanById(id int64, loan *models.Loan) (*mongo.UpdateResult, error){
	iv := bson.M{"_id": id}
	fv := bson.M{"$set": &loan}
	result, err := l.mongoCollection.UpdateOne(l.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (l *Loans)DeleteLoanById(id int64) (*mongo.DeleteResult, error){
	filter := bson.M{"_id": id}
	result, err := l.mongoCollection.DeleteOne(l.ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
