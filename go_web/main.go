package main

import (
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
    "github.com/gin-gonic/gin"
    "fmt"
    docs "go_web/docs"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

var recipes=[] Recipe{}
var recipe Recipe

func init(){
    fmt.Println("At init")
}

func main() {
    // Initialize the Gin router
    r := gin.Default()
    
    r.GET("/",func (c *gin.Context){
        c.String(200,"Connection Ok!")
    })
    r.GET("/getAllRecipes",getAllRecipes)
    r.POST("/addRecipe",addRecipe)
        
    // For Swagger
    // Basepath determines the urls for the api's in the comment annotation
    docs.SwaggerInfo.BasePath = "/"
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    // Run the server on port 8080
    r.Run(":8080")
}

type Recipe struct{
    Id string`json:"id"`
    Name string `json:"name"`
    Country string `json:"country"`
}

// gets all recipes
func getAllRecipes(c *gin.Context){
    c.JSON(200,recipes)
}


// @Summary Adds a recipe
// @Description Adds a new recipe to existing recipes
// @Accept json
// @Produce json
// @Param request body CreateRecipeRequest true "Recipe info"
// @Success 200 {array} Recipe
// @Failure 404 {string} string "Incorrect URL"
// @Failure 400 {string} string "Bad request"
// @Router /addRecipe  [post]
func addRecipe(c *gin.Context){
    if err:=c.ShouldBindJSON(&recipe); err!=nil{
        c.String(400,err.Error())
        return
    }
    
    recipes=append(recipes,recipe)
    fmt.Println(recipes)
    c.String(200,"Added Succesfully")
}

type CreateRecipeRequest struct {
    ID      string    `json:"id"`
    Name    string `json:"name"`
    Country string `json:"country"`
}



