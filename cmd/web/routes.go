package main

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pedidosya/@project_name@/middlewares"

	"github.com/pedidosya/@project_name@/cmd/web/handlers"
	"github.com/pedidosya/@project_name@/engine"
	"github.com/pedidosya/@project_name@/models"
)

type Handler func(cxt *gin.Context, e engine.Spec)

func setupRouter(c *models.Configuration, app newrelic.Application, e engine.Spec) *gin.Engine {
	router = gin.Default()
	//health check route
	router.GET("/", middlewares.NewRelicMonitoring(app), specHandler(e, handlers.HandleHealthCheck))
	router.GET("/health", middlewares.NewRelicMonitoring(app), specHandler(e, handlers.HandleHealthCheck))

	//demo individual route
	router.Group("/v1").GET("/hello/:name", middlewares.NewRelicMonitoring(app), specHandler(e, handlers.HandleHelloWorld))
	// demo grouping routes and adding middlewares by groups instead of individual routes
	v2 := router.Group("/v2")
	v2.Use(middlewares.NewRelicMonitoring(app))
	v2.Use(middlewares.AuthResourceHandler(&c.Auth, []string{}))
	{
		v2.GET("/hello/:name", specHandler(e, handlers.HandleHelloWorld))
		//v2.GET("/hello/:lastname", specHandler(e, handlers.HandleHelloWorld))
	}

	return router
}

func specHandler(e engine.Spec, f Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f(ctx, e)
	}
}
