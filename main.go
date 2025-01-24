package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getUserNameHandler(context *gin.Context) {
    name := context.Param("name")
    context.JSON(
        http.StatusOK,
        gin.H {
            "name": name,
        },
    )
}

func main() {
    engine := gin.Default()

    engine.GET("/user/:name", getUserNameHandler)

    engine.Run()
}
