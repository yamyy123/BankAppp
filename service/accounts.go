package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Acc struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitAccount(collection *mongo.Collection, ctx context.Context) interfaces.Iaccount{
	return &Acc{ctx,collection}
}

func (a *Acc)CreateAccount(account *models.Account)(*mongo.InsertOneResult,error){
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "account_id", Value: 1},{Key: "customer_id", Value: 1}}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := a.mongoCollection.Indexes().CreateMany(a.ctx, indexModel)
	if err != nil {
		return nil,err
	}
	result,err := a.mongoCollection.InsertOne(a.ctx,account)
	if(err!=nil){
		return nil,err
	}
	return result,nil
}


func (a *Acc)GetAccountById(id int64)(*models.Account,error){
	filter := bson.D{{Key: "account_id", Value: id}}
	var account *models.Account
	res := a.mongoCollection.FindOne(a.ctx, filter)
	err := res.Decode(&account)
	if err!=nil{
		return nil,err
	}
	return account,nil
}

func (a *Acc)UpdateAccountById(id int64, account *models.Account)(*mongo.UpdateResult,error){
	iv := bson.M{"account_id": id}
	fv := bson.M{"$set": &account}
	res,err := a.mongoCollection.UpdateOne(a.ctx, iv, fv)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (a *Acc)DeleteAccountById(id int64)(*mongo.DeleteResult,error){
	del := bson.M{"account_id": id}
	res,err := a.mongoCollection.DeleteOne(a.ctx, del)
	if err!=nil{
		return nil,err
	}
	return res,nil
}