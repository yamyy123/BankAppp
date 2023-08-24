package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func AccRoute(router *gin.Engine, controller controllers.AccountController){
	router.POST("/account", controller.CreateAccount)
	router.GET("/account/:id", controller.GetAccountById)
	router.PUT("/account/:id", controller.UpdateAccountById)
	router.DELETE("/account/:id", controller.DeleteAccountById)
}