package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET Opening."})
}

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "POST Opening."})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PUT Opening."})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DELETE Opening."})
}

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET Openings."})
}
