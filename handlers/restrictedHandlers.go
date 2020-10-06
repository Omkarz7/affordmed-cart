package handlers

import (
	"affordmed-cartservice/models"
	"affordmed-cartservice/services"
	"auction-backend/http"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var item models.ItemDetails
		ctx.Bind(&item)
		token := ctx.GetHeader("Authorization")
		url := "http://localhost:3000/users/" + item.UserName + "/validateToken"
		req, err := http.NewRequest("GET", url)
		if err != nil {
			ctx.AbortWithError(401, err)
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			ctx.AbortWithError(401, err)
		}

		var authRes models.AuthResponse
		resByte, err := ioutil.ReadAll(res.Body)
		if err != nil {
			ctx.AbortWithError(401, err)
		}
		err = json.Unmarshal(resByte, &authRes)
		if err != nil {
			ctx.AbortWithError(401, err)
		}

		//some error occured which validating token
		if authRes.Message != "" {
			ctx.AbortWithError(401, error.New(authRes.Message))
		}

		services.AddToCart(authRes)
		ctx.JSONP(200, gin.H{"message": "Item Added to cart"})
	}
}

func RemoveFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
