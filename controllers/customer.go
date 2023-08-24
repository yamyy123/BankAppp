package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{
     TransactionService  interfaces.Icustomer
}


func InitTransController(transactionService interfaces.Icustomer) TransactionController {
    return TransactionController{transactionService}
}

func (t *TransactionController)CreateCustomer(ctx *gin.Context){
    var trans *models.Customer  
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.TransactionService.CreateCustomer(trans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}



func (t *TransactionController)GetCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    val, err := t.TransactionService.GetCustomerById(id1)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *TransactionController)UpdateCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    fv := &models.UpdateModel{}
    if err := ctx.ShouldBindJSON(&fv); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    fmt.Println(fv)
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.TransactionService.UpdateCustomerById(id1,fv)
    if err!=nil{
        fmt.Println("error")
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController)DeleteCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.TransactionService.DeleteCustomerById(id1)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController)GetAllCustomerTransaction(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.TransactionService.GetAllCustomerTransaction(id1)                                                                
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}