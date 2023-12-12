package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

 func getRoute(c *gin.Context){
    fmt.Println(c.Params)
    c.JSON(200,gin.H{
            "Woo":"Hoo",
            "Hooo":"Wooo",
    })
}

func main() {
    // Initialize the Gin router
    r := gin.Default()
    r.GET("/",getRoute)
    r.GET("/:name",getRoute)
    // Run the server on port 8080
    r.Run(":8080")
}

