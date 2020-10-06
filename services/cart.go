package services

import (
	"affordmed-cartservice/models"
)

func AddToCart(item models.ItemDetails) {
	models.Cart = append(models.Cart, item)
}
