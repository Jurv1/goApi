package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"resp": "Yes",
		})
	})

	app.Run(":8080")
}
