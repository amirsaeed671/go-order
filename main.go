package main

import (
	"encoding/json"
	"fmt"
	"golang/utils"
	"io"
	"net/http"
	"os"
)

type ResponseItem struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ResponseItemList struct {
	Items []ResponseItem `json:"items"`
}

func main() {
	var username string

	resp, err := http.Get("https://raw.githubusercontent.com/amirsaeed671/go-order/main/items.json")

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var itemList ResponseItemList

	json.Unmarshal(body, &itemList)

	menu := utils.MakeMenu("1 Star Restaurant")

	for _, val := range itemList.Items {
		itm := utils.MakeItem(val.Id, val.Name, val.Price)
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
	fmt.Print("For balance enquiry press 2\n")
	fmt.Print("For adding balance 3\n\n")
	fmt.Println("=======================")

	fmt.Scan(&screenSelection)

	if screenSelection == 1 {
		order := utils.CreateOrder()
		openOrderMenu(user, menu, &order)
	} else if screenSelection == 2 {
		openUserManagementMenu(user, menu)
	} else if screenSelection == 3 {
		addBalance(user, menu)
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

func addBalance(user *utils.User, menu utils.Menu) {
	user.ShowMyBalance()
	fmt.Println("Enter the amount you want to add to your balance")

	var amount int
	fmt.Scan(&amount)

	user.AddBalance(amount)
	openMenuScreen(user, menu)
}
