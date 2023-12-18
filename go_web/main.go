package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "strconv"
    docs "go_web/docs"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

type Problem struct{
    Name string `json:"name"`
    Tags []string `json:"tags"`
}


var problem Problem

func getData(c *gin.Context){

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

type ProblemResponse struct{
    Tag string
}

// @Summary Return the corresponding tag value
// @Description Return "Person" tag belonging to a particular index
// @ID get-string-by-int
// @Produce  json
// @Param   idx path int true "index value"
// @Success 200 {object} ProblemResponse   
// @Failure 400 {object} string "Bad Req"
// @Router /problem/{idx} [get]
func outData(c *gin.Context){
    var param =c.Param("index")
    num,err:=strconv.Atoi(param)
    if num<0 || num+1>len(problem.Tags) || err!=nil{
        c.String(400,"Error-Bad Req")
        return
    }
    c.JSON(201,gin.H{
        "Tag":problem.Tags[num],
    })
}

func main() {
    // Initialize the Gin router
    r := gin.Default()
    r.GET("/",func (c *gin.Context){
        c.String(200,"Ok!")
    })
    r.POST("/problem",getData)
    r.GET("/problem",func(c *gin.Context){
        if len(problem.Name)==0{
            c.JSON(400,gin.H{
                "Err":"Instance unintialised",
            })
        }
        c.JSON(200,problem.Name)
    })
    r.GET("/problem/:index",outData)
    
    // For Swagger
    // Basepath determines the urls for the api's in the comment annotation
    docs.SwaggerInfo.BasePath = "/"
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    // Run the server on port 8080
    r.Run(":8080")
}

