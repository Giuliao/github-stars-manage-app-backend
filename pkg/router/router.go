package router

import "github.com/gin-gonic/gin"

// InitRouter init router
func InitRouter(router *gin.RouterGroup) {
	initGithubRouter(router)
}
