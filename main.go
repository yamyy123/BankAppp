package main

import (
	"bankDemo/config"
	"bankDemo/constants"
	"bankDemo/controllers"
	"bankDemo/routes"
	"bankDemo/service"
	"context"
	"fmt"
	"log"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)
func initRoutes(){
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database(constants.Dbname).Collection("customer")
	profileService := service.InitCustomer(profileCollection, ctx)
	profileController := controllers.InitTransController(profileService)
	routes.CustRoute(server,profileController)
}

func initAcc(mongoClient *mongo.Client){
	ctx = context.TODO()
	accCollection := mongoClient.Database(constants.Dbname).Collection("account")
	accService := service.InitAccount(accCollection, ctx)
	accController := controllers.InitAccController(accService)
	routes.AccRoute(server,accController)
}

func initBank(mongoClient *mongo.Client){
	ctx = context.TODO()
	bCollection := mongoClient.Database(constants.Dbname).Collection("bank")
	bService := service.InitBank(bCollection, ctx)
	bController := controllers.InitBankController(bService)
	routes.BankRoute(server,bController)
}
func initLoan(mongoClient *mongo.Client){
	ctx = context.TODO()
	lCollection := mongoClient.Database(constants.Dbname).Collection("loan")
	lService := service.InitLoan(lCollection, ctx)
	lController := controllers.InitLoanController(lService)
	routes.LoanRoute(server,lController)
}

func main(){
	server = gin.Default()
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	initAcc(mongoclient)
	initBank(mongoclient)
	initLoan(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}