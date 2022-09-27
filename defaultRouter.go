package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func readItems(c *gin.Context) {
	var im ItemModel

	items, err := im.SelectAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find items"})
	}
	c.IndentedJSON(http.StatusOK, items)
}

func readItemById(c *gin.Context) {
	var im ItemModel
	var err error

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	receivedItem, err := im.SelectById(id)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find object"})
		return
	}
	c.IndentedJSON(http.StatusOK, receivedItem)
	//for _, items := range TestItems {
	//	if items.Id == id {
	//		c.IndentedJSON(http.StatusOK, items)
	//		return
	//	}
	//}
}

func createItem(c *gin.Context) {
	var receivedItem Item
	var im ItemModel

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
		return
	}

	_, err := im.Insert(receivedItem)
	if err != nil {
		fmt.Println("Error", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not create item"})
		return
	}

	c.IndentedJSON(http.StatusCreated, receivedItem)
}

func updateItemById(c *gin.Context) {
	var receivedItem Item

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for i := range TestItems {
		if TestItems[i].Id == id {
			TestItems[i] = receivedItem
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Item updated"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find Item by id"})
}

func deleteItemById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, items := range TestItems {
		if items.Id == id {
			//delete func
			c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted Item by id"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find Item by id"})
}
