package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main/db"
)

func main() {
	var err error

	err = godotenv.Load("env.env")

	router := gin.Default()

	router.GET("/items", readItems)
	router.GET("/items/:id", readItemById)

	router.POST("/items", createItem)

	router.PUT("/items/:id", updateItemById)

	router.DELETE("/items/:id", deleteItemById)

	db.Init()

	err = router.Run("localhost:5050")
	if err != nil {
		fmt.Println("Error in launching router", err)
	}
}
