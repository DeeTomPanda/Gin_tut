package main

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/gin-gonic/gin"
    "os"
    "fmt"
    "context"
    "log"
    docs "go_web/docs"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    handlers "go_web/handlers"
)
// Context to perform timeout tasks
var ctx context.Context
var err error
// MongoDb drivers
var client *mongo.Client
var collection *mongo.Collection
var recipeHandler *handlers.RecipeHandler

var recipes=[] Recipe{}
var recipe Recipe

func init(){
    fmt.Println("At init")
    name:=os.Getenv("username")
    pass:=os.Getenv("password")
    connectionString:=fmt.Sprintf(`mongodb+srv://%s:%s@cluster0.zqcpn82.mongodb.net/`,name,pass)
    // Context is for handling timeouts and go_routines?
    ctx = context.Background()
    client, err = mongo.Connect(ctx,options.Client().ApplyURI(connectionString))
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal("Error pinging MongoDB:", err)
    }

    collection=client.Database("Gin_Tutorial").Collection("Recipes")
    if collection==nil{
        fmt.Println("Databse or Collection ot found")
    }

    recipeHandler=handlers.NewRecipeHandler(ctx,collection)

    log.Println("Connected to MongoDB")
}



func main() {
    // Initialize the Gin router
    r := gin.Default()
    
    r.GET("/",func (c *gin.Context){
        c.String(200,"Connection Ok!")
    })
    r.GET("/getAllRecipes",recipeHandler.GetAllRecipes)
    r.POST("/addRecipe",addRecipe)
    r.POST("/delRecipe",recipeHandler.DelRecipe)
    r.POST("/updateRecipe",recipeHandler.UpdateRecipe)
        
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
    
    // Convert Recipe struct to BSON document using bson.M
    recipeDocument := bson.M{
        "id":      recipe.Id,
        "name":    recipe.Name,
        "country": recipe.Country,
        // Add other fields similarly...
    }

    recipes=append(recipes,recipe)

    _,err:=recipeHandler.Collection.InsertOne(ctx,recipeDocument)
    if err!=nil{
        fmt.Println("Error writing to mongdb")
        c.String(500,"Error in connection",err.Error())
        return
    }

    c.String(200,"Added Succesfully")
}

// For the swagger documentation

type CreateRecipeRequest struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Country string `json:"country"`
}



