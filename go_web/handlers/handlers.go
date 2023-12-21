package handlers

import (
	"github.com/gin-gonic/gin"
	models "go_web/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"fmt"
)

type RecipeHandler struct{
	Ctx context.Context
	Collection *mongo.Collection
}


// The Main function called in init() that initializes and creates an instance 
// of the context and collection
func NewRecipeHandler (ctx context.Context,collection *mongo.Collection) *RecipeHandler{
	return &RecipeHandler{
		Ctx:ctx,
		Collection:collection,
	}
}

// This is a receiver function, i.e an associated function for structs in rust
func (handler *RecipeHandler) GetAllRecipes(c *gin.Context){
	cur,err:=handler.Collection.Find(handler.Ctx,bson.M{})
	if err!=nil{
		fmt.Println("Error when retrieving collection")
		c.JSON(400,gin.H{
			"message":"Err when retrieving collection",
		})
		return
	}

	recipes:=make([]models.Recipe,0)

	for cur.Next(handler.Ctx){
		var recipe models.Recipe
		cur.Decode(&recipe)
		recipes=append(recipes,recipe)
	}

	c.JSON(200,recipes)
}


func (handler *RecipeHandler) DelRecipe(c *gin.Context){
    // interface{} is atype to hold data of unknown type, like _ in Rust
    var reqBody map[string]interface{}

    if err:=c.ShouldBindJSON(&reqBody); err!=nil{
        c.String(400,err.Error())
        return
    }

    id,ok:=reqBody["id"].(string)
    if !ok{
        c.String(400,"Error while retrieving id")
        return
    }

    filter:=bson.M{"id":id}
    _,err:=handler.Collection.DeleteOne(handler.Ctx,filter)

    if err!=nil{
        c.JSON(400,"Err in deletion")
        return
    }

    c.String(200,"Successfully deleted id")
}


func (recipe *RecipeHandler) UpdateRecipe(c *gin.Context){
    var updatedRecipe models.Recipe

    if err:=c.ShouldBindJSON(&updatedRecipe);err!=nil{
        c.String(400,"Data sent is malformed")
        return
    }

    filter:=bson.M{"id":updatedRecipe.Id}
    // the following is important for upsert ops
    update:=bson.M{"$set":bson.M{"name":updatedRecipe.Name,"country":updatedRecipe.Country}}
    opts:=options.Update().SetUpsert(true)

    if _,err:=recipe.Collection.UpdateOne(recipe.Ctx,filter,update,opts);err!=nil{
        c.JSON(500,"ERROR while updating document")
        return
    }

    c.JSON(200,gin.H{"message":"Success"})
}