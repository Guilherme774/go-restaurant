package controllers

import (
	"GoRestaurant/models"
	"fmt"
	"strconv"
)

func ShowFoods() {
	foods := []models.Item{
		{Description: "Burger", Price: 8.99},
		{Description: "Fries", Price: 3.99},
		{Description: "Fish & Chips", Price: 1.99},
	}

	fmt.Print("What do you want to eat?\n")

	for i := 0; i < len(foods); i++ {
		index := strconv.Itoa(i + 1)

		price := fmt.Sprintf("%.2f", foods[i].Price)

		fmt.Println("[" + index + "] " + foods[i].Description + " $" + price)
	}
}
