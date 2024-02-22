package main

import (
	"GoRestaurant/controllers"
	"GoRestaurant/models"
	"fmt"
)

func main() {
	var totalOrder float32 = 0
	var itemsOrder []models.Item
	itemSelected := 0

	foods := []models.Item{
		{Description: "Burger", Price: 8.99},
		{Description: "Fries", Price: 3.99},
		{Description: "Fish & Chips", Price: 1.99},
	}
	drinks := []models.Item{
		{Description: "Coffee", Price: 2.99},
		{Description: "Orange Juice", Price: 4.99},
		{Description: "Water", Price: 0.99},
	}
	candies := []models.Item{
		{Description: "Chocolate cake", Price: 7.99},
		{Description: "Cupcake", Price: 4.99},
		{Description: "Strawberry Pie", Price: 9.99},
	}

	fmt.Println("Welcome to the Go's Restaurant!")

	controllers.ShowFoods()

	fmt.Println("\nSelect an item:")
	fmt.Scanln(&itemSelected)

	itemsOrder = append(itemsOrder, foods[itemSelected-1])
	totalOrder += foods[itemSelected-1].Price

	fmt.Println("\nNice choice!")

	controllers.ShowDrinks()

	fmt.Println("\nSelect an item:")
	fmt.Scanln(&itemSelected)

	itemsOrder = append(itemsOrder, drinks[itemSelected-1])
	totalOrder += drinks[itemSelected-1].Price

	fmt.Println("\nThats great!")

	controllers.ShowDesserts()

	fmt.Println("\nSelect an item:")
	fmt.Scanln(&itemSelected)

	itemsOrder = append(itemsOrder, candies[itemSelected-1])
	totalOrder += candies[itemSelected-1].Price

	controllers.ShowOrderDetails(itemsOrder, totalOrder)

	fmt.Println("\nThank you for dininging with us!")
	fmt.Println("See ya!")
}
