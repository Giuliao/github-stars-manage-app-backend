package router

import (
	"github.com/gin-gonic/gin"

	v1 "github-stars/pkg/api/v1"
)

func initGithubRouter(router *gin.RouterGroup) {
	router.GET("/auth", v1.Auth)
}
