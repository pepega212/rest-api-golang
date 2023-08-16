package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mod/configs"
	"go.mod/models"
)

var posts = []models.Post{}

func CreatePostController(c echo.Context) error {
	// binding data
	post := models.Post{}
	c.Bind(&post)

	result := configs.DB.Create(&post)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert post data",
			Data:    nil,
		})
	}

	if len(posts) < 1 {
		post.Id = 1
	} else {
		newId := posts[len(posts)-1].Id + 1
		post.Id = newId
	}

	// append new data to slice posts
	posts = append(posts, post)

	// render JSON response with success message and the new data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create post",
		"post":     post,
	})
}

// get post by id
func GetPostController(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))

	post := models.Post{}

	result := configs.DB.First(&post, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status:  false,
			Message: "post Not Found",
			Data:    nil,
		})
	}

	// if post data is found, return it
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success post found",
		"post":     post,
	})
}

// get all posts
func GetPostsController(c echo.Context) error {
	result := configs.DB.Find(&posts)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to get posts data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all posts",
		"posts":    posts,
	})
}

// update post by id
func UpdatePostController(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))

	// binding data
	input := models.Post{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid input data",
			Data:    nil,
		})
	}

	// find data and change the value, then render JSON response
	for i, v := range posts {
		if v.Id == id {
			posts[i].Title = input.Title
			posts[i].Content = input.Content
			result := configs.DB.Save(&posts[i])
			if result.Error != nil {
				return c.JSON(http.StatusInternalServerError, models.BaseResponse{
					Status:  false,
					Message: "Failed to update post",
					Data:    nil,
				})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success updated post",
				"post":     posts[i],
			})
		}
	}

	// when the data does not exist render JSON response with not found message
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "Post Not Found",
	})
}

// delete posts by id
func DeletePostController(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))

	// looking for data in index of slice posts, then delete element in slice posts and render JSON response
	for i, v := range posts {
		if v.Id == id {
			result := configs.DB.Delete(&posts[i], id)
			if result.Error != nil {
				return c.JSON(http.StatusInternalServerError, models.BaseResponse{
					Status:  false,
					Message: "Failed to delete post",
					Data:    nil,
				})
			}
			posts = append(posts[:i], posts[i+1:]...)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "delete success",
			})
		}
	}

	// when the data does not exist render JSON response with not found message
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "Post Not Found",
	})
}

func GetPostsByUserController(c echo.Context) error {
	// catch path param from url
	userId, _ := strconv.Atoi(c.Param("userId"))

	// find posts by user ID
	userPosts := []models.Post{}
	result := configs.DB.Where("user_id = ?", userId).Find(&userPosts)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to get user's posts",
			Data:    nil,
		})
	}

	// Check if any posts were found for the user ID
	if len(userPosts) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success get posts by user ID",
			"posts":    userPosts,
		})
	}

	// If no posts were found, return the appropriate response
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user post not found",
	})
}
