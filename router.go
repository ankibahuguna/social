package main

import (
	"github.com/ankibahuguna/social/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (app *App) InitRouter() {
	nonAuth := app.Echo.Group("/api/v1")

	nonAuth.GET("/users/:id", echo.HandlerFunc(controller.GetUser))
	nonAuth.POST("/users", echo.HandlerFunc(controller.SaveUser))
	nonAuth.POST("/users/login", echo.HandlerFunc(controller.Login))

	nonAuth.GET("/posts", echo.HandlerFunc(controller.GetAllPosts))
	nonAuth.GET("/posts/:id", echo.HandlerFunc(controller.GetSinglePost))

	auth := app.Group("/api/v1")

	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("thisIsASafeSe2374823478#$$%$%^key"),
		ContextKey: "user",
	}))

	auth.GET("/users", echo.HandlerFunc(controller.GetUsers))
	auth.PUT("/users/password", echo.HandlerFunc(controller.ChangePassword))
	auth.POST("/users/password", echo.HandlerFunc(controller.ResetPassword))
	auth.POST("/users/:id/follow", echo.HandlerFunc(controller.FollowUser))
	auth.DELETE("/users/:id/unfollow", echo.HandlerFunc(controller.UnFollowUser))
	auth.GET("/users/:id/followers", echo.HandlerFunc(controller.GetUserFollowers))
	auth.GET("/profile", echo.HandlerFunc(controller.GetUserProfile))

	auth.POST("/posts", echo.HandlerFunc(controller.CreateNewPost))
	auth.DELETE("/posts/:id", echo.HandlerFunc(controller.DeleteSinglePost))

	auth.POST("/posts/:id/comments", echo.HandlerFunc(controller.PostComment))
	auth.PUT("/posts/:id/comments", echo.HandlerFunc(controller.EditComment))
	auth.DELETE("/posts/:id/comments", echo.HandlerFunc(controller.DeleteComment))
	auth.GET("/posts/:id/comments", echo.HandlerFunc(controller.GetComments))

	auth.POST("/posts/:id/likes", echo.HandlerFunc(controller.LikePost))
	auth.DELETE("/posts/:id/likes", echo.HandlerFunc(controller.UnlikePost))
	auth.GET("/posts/:id/likes", echo.HandlerFunc(controller.GetLikes))

}
