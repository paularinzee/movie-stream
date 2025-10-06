package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/paularinzee/server/movie-stream-api/controllers"
	"github.com/paularinzee/server/movie-stream-api/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func MovieRoutes(router *gin.Engine, client *mongo.Client) {

	router.GET("/movies", controller.GetMovies(client))

	router.GET("/genres", controller.GetGenres(client))

	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.POST("/addmovie", controller.AddMovie(client))
	router.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate(client))
}
