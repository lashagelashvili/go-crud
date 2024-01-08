package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data from request
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return post

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// return them
	c.JSON(200, gin.H{
		"post": post,
	})
}
