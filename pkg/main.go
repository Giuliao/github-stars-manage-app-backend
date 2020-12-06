package main

import (
	"github-stars/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1GroupRouter := r.Group("/api/v1")
	router.InitRouter(v1GroupRouter)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error_message": "no api registerred"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
