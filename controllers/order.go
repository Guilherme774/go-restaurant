package controllers

import (
	"GoRestaurant/models"
	"fmt"
)

func ShowOrderDetails(itemsSelectedOrder []models.Item, totalOrder float32) {
	fmt.Println("\nHere is your order:")

	fmt.Println("Dinner:")
	fmt.Println("- " + itemsSelectedOrder[0].Description + " $" + fmt.Sprintf("%.2f", itemsSelectedOrder[0].Price))

	fmt.Println("Drink:")
	fmt.Println("- " + itemsSelectedOrder[1].Description + " $" + fmt.Sprintf("%.2f", itemsSelectedOrder[1].Price))

	fmt.Println("Dessert:")
	fmt.Println("- " + itemsSelectedOrder[2].Description + " $" + fmt.Sprintf("%.2f", itemsSelectedOrder[2].Price))

	fmt.Println("\nYour total is $" + fmt.Sprintf("%.2f", totalOrder))
}
