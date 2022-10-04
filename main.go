package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"main/controllers"
	"main/db"
)

func main() {
	var err error

	err = godotenv.Load("env.env")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/items", controllers.ReadItems)
	router.GET("/items/:id", controllers.ReadItemById)

	router.POST("/items", controllers.CreateItem)

	router.PUT("/items/:id", controllers.UpdateItemById)

	router.DELETE("/items/:id", controllers.DeleteItemById)

	db.Init()

	err = router.Run("localhost:5050")
	if err != nil {
		fmt.Println("Error in launching router", err)
	}
}
