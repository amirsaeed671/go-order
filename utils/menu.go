package utils

type Menu struct {
	RestaurantName string
	Items          []Item
}

func (m *Menu) AddItems(item Item) {
	newItems := append(m.Items, item)

	m.Items = newItems
}

func MakeMenu(restaurantName string) Menu {
	return Menu{RestaurantName: restaurantName}
}
