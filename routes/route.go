package routes

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mod/controllers"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.POST("/login", controllers.LoginController)

	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY_JWT"))))

	auth.GET("/users", controllers.GetUsersController)                     //get all users
	auth.GET("/users/:id", controllers.GetUserController)                  //get user by id
	auth.GET("/posts", controllers.GetPostsController)                     //get all posts
	auth.GET("/posts/:id", controllers.GetPostController)                  //get post by id
	auth.GET("/users/:userId/posts", controllers.GetPostsByUserController) //get post by user id
	auth.POST("/posts", controllers.CreatePostController)                  // Create post
	auth.POST("/users", controllers.CreateUserController)                  //Create new user
	auth.PUT("/users/:id", controllers.UpdateUserController)               //update user
	auth.PUT("posts/:id", controllers.UpdatePostController)                //update post
	auth.DELETE("/posts/:id", controllers.DeletePostController)            //delete post
}
