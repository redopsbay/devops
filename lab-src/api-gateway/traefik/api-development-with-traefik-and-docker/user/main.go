package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"id": 1, "username": "andres"})
  })
  r.Run(":8000")
}