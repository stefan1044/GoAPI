package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
	"strconv"
)

func CreateItem(c *gin.Context) {
	var im models.ItemModel
	var receivedItem models.Item

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

func DeleteItemById(c *gin.Context) {
	//var im ItemModel
	//id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find Item by id"})
}

func ReadItemById(c *gin.Context) {
	var im models.ItemModel
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
}

func ReadItems(c *gin.Context) {
	var im models.ItemModel

	items, err := im.SelectAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find items"})
	}
	c.IndentedJSON(http.StatusOK, items)
}

func UpdateItemById(c *gin.Context) {
	var im models.ItemModel
	var receivedItem models.Item

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = im.UpdateById(id, receivedItem)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not update item"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
