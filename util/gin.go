package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data, "error": nil})
}
func Error(c *gin.Context, statusCode int, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{"success": false, "data": nil, "error": err.Error()})
		return
	}

	c.JSON(statusCode, gin.H{"success": false, "data": nil, "error": nil})
}
