package controllers

import (
	"GoRestaurant/models"
	"fmt"
	"strconv"
)

func ShowDesserts() {
	desserts := []models.Item{
		{Description: "Chocolate cake", Price: 7.99},
		{Description: "Cupcake", Price: 4.99},
		{Description: "Strawberry Pie", Price: 9.99},
	}

	fmt.Print("What do you want for dessert?\n")

	for i := 0; i < len(desserts); i++ {
		index := strconv.Itoa(i + 1)

		Price := fmt.Sprintf("%.2f", desserts[i].Price)

		fmt.Println("[" + index + "] " + desserts[i].Description + " $" + Price)
	}

}
