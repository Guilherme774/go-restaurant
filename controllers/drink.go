package controllers

import (
	"GoRestaurant/models"
	"fmt"
	"strconv"
)

func ShowDrinks() {
	drinks := []models.Item{
		{Description: "Coffee", Price: 2.99},
		{Description: "Orange Juice", Price: 4.99},
		{Description: "Water", Price: 0.99},
	}

	fmt.Print("What do you want to drink?\n")

	for i := 0; i < len(drinks); i++ {
		index := strconv.Itoa(i + 1)

		price := fmt.Sprintf("%.2f", drinks[i].Price)

		fmt.Println("[" + index + "] " + drinks[i].Description + " $" + price)
	}
}
