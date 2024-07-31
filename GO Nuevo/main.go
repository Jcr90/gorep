package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *gorm.DB

func initDB() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})
}

func main() {
	initDB()
	userRepo := NewUserRepository(db)
	handler := NewHandler(userRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rutas públicas
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)

	// Rutas protegidas
	r := e.Group("")
	r.Use(AuthMiddleware)
	r.POST("/users", handler.CreateUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
