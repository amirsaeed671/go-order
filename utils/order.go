package utils

import (
	"fmt"
)

type Order struct {
	Status string
	Items  []Item
	Total  int
}

func (o *Order) AddItem(item Item, user *User) {
	if user.Balance < item.Price {
		fmt.Println("=================================================")
		fmt.Println("Your Balance is insufficient for adding this item")
		fmt.Println("=================================================")

		return
	}

	newItems := append(o.Items, item)

	o.Items = newItems
	o.Total += item.Price
}

func (o *Order) PlaceOrder(user *User) {
	o.Status = "Completed"

	o.PrintReceipt()
	user.Balance -= o.Total
}

func (o *Order) PrintReceipt() {
	fmt.Println("=========================")
	fmt.Println("Here is your Order receipt")
	fmt.Println("=========================")

	fmt.Println("List of items ordered")

	for _, item := range o.Items {
		fmt.Printf("Item name: %+v, Price: %d\n", item.Name, item.Price)
	}

	fmt.Println("Your Order Status : ", o.Status)
	fmt.Println("Your Order Total : ", o.Total)

}

func CreateOrder() Order {
	return Order{Status: "Pending", Total: 0}
}
