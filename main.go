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

func getUserNameActionHandler(context *gin.Context) {
    name := context.Param("name")
    action := context.Param("action")
    context.JSON(
        http.StatusOK,
        gin.H {
            "name": name,
            "action": action,
        },
    )
}

func main() {
    engine := gin.Default()

    engine.GET("/user/:name", getUserNameHandler)
    engine.GET("/user/:name/:action", getUserNameActionHandler)

    engine.Run()
}
