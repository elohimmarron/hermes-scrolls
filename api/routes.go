package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{})
    })

    return r
}

func StartServer() {
    r := SetupRouter()
    r.Run(":8080")
}