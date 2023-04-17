package handlers

import (
	"github.com/pedidosya/@project_name@/engine"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pedidosya/@project_name@/models"
)

var AppConfig *models.Configuration

type Handler func(ctx *gin.Context, e engine.Spec)

func initialize() {
	AppConfig = &models.Configuration{
		Vault:    models.ExternalServiceConfig{},
		NewRelic: &newrelic.Config{},
		App:      models.AppConfig{},
		Auth:     models.ExternalServiceConfig{},
	}
}

// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
	initialize()
	gin.SetMode(gin.TestMode)
	// Run the other tests
	os.Exit(m.Run())
}

func getRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func getEngine() engine.Spec {
	return engine.New(AppConfig)
}

func specHandler(e engine.Spec, f Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f(ctx, e)
	}
}