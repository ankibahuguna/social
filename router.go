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
	auth.GET("/profile", echo.HandlerFunc(controller.GetUserProfile))
	auth.POST("/posts", echo.HandlerFunc(controller.CreateNewPost))

}
