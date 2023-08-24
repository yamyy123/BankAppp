package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Cust struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}


func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer{
	return &Cust{ctx,collection}
}
func(c *Cust) CreateCustomer(user *models.Customer)(*mongo.InsertOneResult,error){
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "account_id", Value: 1},{Key: "customer_id", Value: 1}}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := c.mongoCollection.Indexes().CreateMany(c.ctx, indexModel)
	if err != nil {
		return nil,err
	}
	user.Transaction[0].Date = time.Now()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password),7)
	user.Password = string(hashedPassword)
	res,err := c.mongoCollection.InsertOne(c.ctx, &user)
	if err!=nil{
		if mongo.IsDuplicateKeyError(err){
			log.Fatal("Duplicate key error")
		}
		return nil,err
	}
	
	return res,nil
}


func(c *Cust) GetCustomerById(id int64) (*models.Customer, error) {
	filter := bson.D{{Key: "customer_id", Value: id}}
	var customer *models.Customer
	res := c.mongoCollection.FindOne(c.ctx, filter)
	err := res.Decode(&customer)
	if err!=nil{
		return nil,err
	}
	return customer,nil
}

func(c *Cust) UpdateCustomerById(id int64, n *models.UpdateModel) (*mongo.UpdateResult, error){
	iv := bson.M{"customer_id": id}
	if n.Topic == "password"{
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(string(n.FinalValue.(string))),8)
		n.FinalValue = string(hashedPassword)
	}
	if reflect.TypeOf(n.FinalValue).String() == "float64"{
		n.FinalValue = int64(n.FinalValue.(float64))
	}
	fv := bson.M{"$set": bson.M{n.Topic: n.FinalValue}}
	res,err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Cust) DeleteCustomerById(id int64) (*mongo.DeleteResult, error){
	del := bson.M{"customer_id": id}
	res,err := c.mongoCollection.DeleteOne(c.ctx, del)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Cust)GetAllCustomerTransaction(id int64)(*[]models.CustTransaction,error){
	filter := bson.D{{Key: "customer_id", Value: id}}
	var customer *models.Customer
	res := c.mongoCollection.FindOne(c.ctx, filter)
	err := res.Decode(&customer)
	if err!=nil{
		return nil,err
	}
	return &customer.Transaction,nil
}