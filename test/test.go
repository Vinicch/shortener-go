package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/vinicch/shortener-go/handlers"
	"github.com/vinicch/shortener-go/logging"
	"github.com/vinicch/shortener-go/repository"
)

func setup() *gin.Engine {
	logging.Setup()
	godotenv.Load("../.env")

	urlFunctions := repository.MakeURLFunctions()
	router := gin.Default()

	router.POST("/create", handlers.Create(urlFunctions.CreateURL, urlFunctions.DoesAliasExist))
	router.GET("/url/:alias", handlers.Retrieve(urlFunctions.GetURL, urlFunctions.UpdateURL))
	router.GET("/most-visited", handlers.MostVisited(urlFunctions.GetMostVisited))

	return router
}

func TestCreate(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create?url=http://hostname.com/long/url/path", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestCreateConflict(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create?url=http://hostname.com/long/url/path&CUSTOM_ALIAS=test", nil)

	// Runs twice to guarantee conflict
	router.ServeHTTP(&httptest.ResponseRecorder{}, req)
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusConflict, recorder.Code)
}

func TestRetrieve(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/url/test", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusMovedPermanently, recorder.Code)
}

func TestRetrieveNotFound(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/url/qwertyuiop", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestMostVisited(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/most-visited", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
