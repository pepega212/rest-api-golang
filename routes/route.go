package routes

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
	"go.mod/controllers"
)

func InitRoute(e *echo.Echo) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"}, // Replace with your local server's origin
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	})
	e.Use(middleware.Logger())
	e.Use(echo.WrapMiddleware(c.Handler))
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
