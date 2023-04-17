package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pedidosya/@project_name@/helpers"
	"github.com/pedidosya/@project_name@/models"
	"github.com/pedidosya/peya-go/common"
	"github.com/pedidosya/peya-go/logs"
	"net/http"
)

func AuthResourceHandler(cfg *models.ExternalServiceConfig, resources []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if len(token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":     "FORBIDDEN",
				"messages": "no token was provided",
			})
			return
		}
		cred := &common.Credentials{}
		client, err := helpers.NewRequest("GET", cfg.URL, "", cred)
		if err != nil {
			logs.Errorf("middleware_auth GetCredentials: %v", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		err = client.Do(ctx, &cred)

		if cred == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":     "FORBIDDEN",
				"messages": "invalid token",
			})
			return
		}

		if err != nil {
			logs.Errorf("middleware_auth GetCredentials: %v", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		if len(resources) > 0 {
			if ok := validateTokenResources(cred, resources); !ok {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"code":     "FORBIDDEN",
					"messages": "not authorized",
				})
				return
			}
		}
		ctx.Set(common.CredentialsInContext, cred)
		ctx.Next()
	}
}

func validateTokenResources(token *common.Credentials, resources []string) bool {
	res := map[string]bool{}
	if len(resources) == 0 {
		return true
	}
	for _, r := range token.Resources {
		res[r] = true
	}
	for _, r := range resources {
		if _, ok := res[r]; !ok {
			return false
		}
	}
	return true
}
