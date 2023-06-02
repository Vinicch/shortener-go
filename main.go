package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vinicch/shortener-go/adapters/logging"
	"github.com/vinicch/shortener-go/adapters/repository"
	"github.com/vinicch/shortener-go/adapters/web"
)

func main() {
	logging.Setup()

	urlFunctions := repository.MakeURLFunctions()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CLIENT_URL")},
	}))

	router.POST("/create", web.Create(urlFunctions.CreateURL, urlFunctions.DoesAliasExist))
	router.GET("/url/:alias", web.Retrieve(urlFunctions.GetURL, urlFunctions.UpdateURL))
	router.GET("/most-visited", web.MostVisited(urlFunctions.GetMostVisited))

	router.Run()
}
