package main

import (
	"github.com/ankibahuguna/social/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	*echo.Echo
	db *gorm.DB
}

var config = struct {
	Host string `default:"localhost"`
	Port string `default:"3000"`
}{}

func main() {
	// Load configuration
	configor.Load(&config)

	app := &App{echo.New(), db.New()}

	app.Debug = true

	// Routing setup
	app.InitRouter()

	defer app.db.Close()

	// Wrap db pointer into echo.context as a middleware
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			context.Set("db", app.db)
			return next(context)
		}
	})
	app.Echo.Use(middleware.Logger())

	// launch
	app.Logger.Fatal(app.Start(config.Host + ":" + config.Port))
}
