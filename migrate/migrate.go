package main

import (
	"github.com/arashzich/go-jwt/initializers"
	"github.com/arashzich/go-jwt/models"
)

func init() {
	initializers.LoadEbvVariables()
	initializers.ConnectToDb()

}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Post{})

}
