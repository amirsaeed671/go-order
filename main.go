package main

import (
	"fmt"
	"golang/utils"
	"os"
)

func main() {
	var username string
	itemsList := map[string][]int{"Chocolate": {1, 50}, "Ice cream": {2, 200}, "Burger": {3, 500}, "Pizza": {4, 350}}
	menu := utils.MakeMenu("1 Star Restaurant")

	for name, val := range itemsList {
		itm := utils.MakeItem(val[0], name, val[1])
		menu.AddItems(itm)
	}

	fmt.Println("Register yourself")
	fmt.Print("Please Enter your name: ")
	fmt.Scanf("%s", &username)

	user := utils.MakeUser(username, 100)

	fmt.Printf("=========================================\n")
	fmt.Printf("Welcome to the %+v, %+v\n", menu.RestaurantName, username)
	fmt.Printf("=========================================\n\n")
	fmt.Print("The menu is as below: \n\n")
	for _, item := range menu.Items {
		item.LogItem()
	}

	openMenuScreen(&user, menu)

}

func openMenuScreen(user *utils.User, menu utils.Menu) {
	var screenSelection int

	fmt.Println("=======================")
	fmt.Print("\nFor exit press 0\n")
	fmt.Print("For order press 1\n")
	fmt.Print("For balance enquiry press 2\n\n")
	fmt.Println("=======================")

	fmt.Scan(&screenSelection)

	if screenSelection == 1 {
		order := utils.CreateOrder()
		openOrderMenu(user, menu, &order)
	} else if screenSelection == 2 {
		openUserManagementMenu(user, menu)
	} else if screenSelection == 0 {
		fmt.Println("Exiting the program. Goodbye!")
		os.Exit(0)
	} else {
		fmt.Println("Please choose one of the given options: 0, 1, 2")
		openMenuScreen(user, menu)
	}
}

func openOrderMenu(user *utils.User, menu utils.Menu, order *utils.Order) {
	fmt.Println("Please enter the id of the item to order, once done write '0' and press enter")

	var option int
	fmt.Scan(&option)

	fmt.Println(option, "************")

	if option == 0 {
		order.PlaceOrder(user)
		openMenuScreen(user, menu)
		return
	} else {
		itemFound := false
		for _, itm := range menu.Items {
			if itm.Id == option {
				order.AddItem(itm, user)
				itemFound = true
				openOrderMenu(user, menu, order)
			}
		}

		if !itemFound {
			fmt.Println("The item you entered does not exist")
			openOrderMenu(user, menu, order)
		}

	}

	openMenuScreen(user, menu)
}

func openUserManagementMenu(user *utils.User, menu utils.Menu) {
	user.ShowMyBalance()
	openMenuScreen(user, menu)
}
