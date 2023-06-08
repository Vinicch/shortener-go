package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vinicch/shortener-go/handlers"
	"github.com/vinicch/shortener-go/logging"
	"github.com/vinicch/shortener-go/repository"
)

func main() {
	logging.Setup()

	urlFunctions := repository.MakeURLFunctions()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CLIENT_URL")},
	}))

	router.POST("/create", handlers.Create(urlFunctions.CreateURL, urlFunctions.DoesAliasExist))
	router.GET("/url/:alias", handlers.Retrieve(urlFunctions.GetURL, urlFunctions.UpdateURL))
	router.GET("/most-visited", handlers.MostVisited(urlFunctions.GetMostVisited))

	router.Run()
}
