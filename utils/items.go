package utils

import (
	"fmt"
)

type Item struct {
	Id    int
	Name  string
	Price int
}

func (i *Item) LogItem() {
	fmt.Println("Id of the item ", i.Id)
	fmt.Println("Name of the item ", i.Name)
	fmt.Println("Price of the item ", i.Price)
}

func MakeItem(id int, name string, price int) Item {
	return Item{Id: id, Name: name, Price: price}
}
