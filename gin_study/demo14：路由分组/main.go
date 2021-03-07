package main

import "github.com/gin-gonic/gin"

func loginEndpoint(c *gin.Context) {
	c.String(200, "login")
}

func submitEndpoint(c *gin.Context) {
	c.String(200, "submit")
}

func readEndpoint(c *gin.Context) {
	c.String(200, "read")
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}
