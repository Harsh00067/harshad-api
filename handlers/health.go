package handlers

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "running it",
	})
}

//when user add some endpoint , for that endpoint handler is there , so it runs when it is written in browser
