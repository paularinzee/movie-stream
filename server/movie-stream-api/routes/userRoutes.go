package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/paularinzee/server/movie-stream-api/controllers"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UserRoutes(router *gin.Engine, client *mongo.Client) {
	router.POST("/register", controller.RegisterUser(client))
	router.POST("/login", controller.LoginUser(client))
	router.POST("/logout", controller.LogoutHandler(client))
	router.POST("/refresh", controller.RefreshTokenHandler(client))
}
