package utils

import "fmt"

type User struct {
	Name    string
	Balance int
	Orders  []Order
}

func (u *User) AddBalance(amount int) {
	u.Balance += amount

	u.ShowMyBalance()
}

func (u *User) ShowMyBalance() {
	fmt.Println("This is your current balance: ", u.Balance)
}

func MakeUser(name string, balance int) User {
	return User{
		Name:    name,
		Balance: balance,
	}
}
