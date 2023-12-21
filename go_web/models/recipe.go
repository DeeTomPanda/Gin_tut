package models

// Basically defines how it is to be serialised 
// when being binded to JSON or BSON
type Recipe struct{
	Id string `bson.M:"id" json:"id"`
	Name string `bson.M:"name" json:"name"`
	Country string `bson.M:"country" json:"country"`
}