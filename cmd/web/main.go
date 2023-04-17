package main

import (
	"fmt"
	"github.com/pedidosya/@project_name@/middlewares"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pedidosya/@project_name@/config"
	"github.com/pedidosya/@project_name@/engine"
)

var env string

var router *gin.Engine

func init() {
	env = getEnv()
	if env == "live" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	startServer(env)
}

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "dev"
	}

	return env
}

func startServer(env string) {
	config := config.LoadConfigurations(env)

	engine := engine.New(config)

	router := setupRouter(config, middlewares.SetupNewRelic(config), engine)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.App.AppPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
