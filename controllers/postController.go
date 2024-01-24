package controllers

import (
	"github.com/arashzich/go-jwt/initializers"
	"github.com/arashzich/go-jwt/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data off req body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {

	// Get the posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond with the posts

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {

	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond with the posts

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {

	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Get data off req body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	c.Bind(&body)

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with the posts

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {

	// Get id off url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with the posts

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
