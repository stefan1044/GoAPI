package main

import (
	"main/db"
	"strconv"
)

type Item struct {
	Id        int64  `json:"id"`
	Action    string `json:"action"`
	Completed bool   `json:"completed"`
}

type ItemModel struct{}

var TestItems = []Item{
	{Id: 1, Action: "Go home", Completed: false},
	{Id: 2, Action: "Eat dinner", Completed: false},
	{Id: 3, Action: "Sleep", Completed: true},
}

func (i ItemModel) Insert(item Item) (itemId int64, err error) {
	var boolString string
	switch {
	case item.Completed:
		boolString = "true"
	default:
		boolString = "false"
	}

	err = db.GetDb().QueryRow("INSERT INTO `items` (`Id`,`Action`,`Completed`)" +
		" VALUES (" + strconv.Itoa(int(item.Id)) + "," + "\"" + item.Action +
		"\"" + "," + boolString + ")").Err()

	//fmt.Println("INSERT INTO `items` (`Id`,`Action`,`Completed`)" +
	//	" VALUES (" + strconv.Itoa(int(item.Id)) + "," + "\"" + string(item.Action) +
	//	"\"" + "," + boolString + ")")

	return
}

func (i ItemModel) SelectById(id int64) (item Item, err error) {
	err = db.GetDb().QueryRow("SELECT Id,Action,Completed FROM items WHERE items.Id = "+
		strconv.Itoa(int(id))).Scan(&item.Id, &item.Action, &item.Completed)
	//fmt.Println("SELECT Id,Action,Completed FROM items WHERE items.Id = "+
	//		strconv.Itoa(int(id)))
	return
}

func (i ItemModel) SelectAll() (items []Item, err error) {
	rows, err := db.GetDb().Query("SELECT Id,Action,Completed FROM items")
	defer rows.Close()
	if err != nil {
		return
	}

	for i := 0; rows.Next(); i++ {
		var temp Item
		if errRow := rows.Scan(&temp.Id, &temp.Action, &temp.Completed); errRow != nil {
			var zero []Item
			return zero, err
		}
		items = append(items, temp)
	}

	return
}
