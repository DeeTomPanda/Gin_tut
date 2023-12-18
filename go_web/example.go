package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

type Problem struct{
    Name string `json:"name"`
    Tags []string `json:"tags"`
}

func getData(c *gin.Context){
    
    var problem Problem
    if err:=c.ShouldBindJSON(&problem); err!=nil{
        fmt.Println("Error while parsing")
        c.JSON(404,gin.H{
            "Error":err.Error(),
        })
        return
    }
    fmt.Println(problem.Tags[0])
    c.JSON(200,gin.H{
        "status":"ok",
    })

}

func main() {
    // Initialize the Gin router
    r := gin.Default()
    r.GET("/",func (c *gin.Context){
        c.String(200,"Ok!")
    })
    r.POST("/problem",getData)
    // Run the server on port 8080
    r.Run(":8080")
}

