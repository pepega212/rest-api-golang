package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mod/configs"
	"go.mod/models"
)

var users = []models.User{}

// get all users
func GetUsersController(c echo.Context) error {
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to get users data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := models.User{}
	c.Bind(&user)

	result := configs.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert user data",
			Data:    nil,
		})
	}

	// if length of slice users less then 1 id = 1, else id +1 from last id in data slice users
	if len(users) < 1 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	// append new data to slice users
	users = append(users, user)

	// render JSON response with success message and the new data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}

	result := configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status:  false,
			Message: "User Not Found",
			Data:    nil,
		})
	}

	// if user data is found, return it
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success user found",
		"user":     user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))

	// binding data
	input := models.User{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid input data",
			Data:    nil,
		})
	}

	// find data and change the value, then render JSON response
	for i, v := range users {
		if v.Id == id {
			users[i].Name = input.Name
			users[i].Email = input.Email
			users[i].Password = input.Password
			users[i].Address = input.Address
			result := configs.DB.Save(&users[i])
			if result.Error != nil {
				return c.JSON(http.StatusInternalServerError, models.BaseResponse{
					Status:  false,
					Message: "Failed to update user",
					Data:    nil,
				})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success updated user",
				"user":     users[i],
			})
		}
	}

	// when the data does not exist render JSON response with not found message
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "User Not Found",
	})
}
