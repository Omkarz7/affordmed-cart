package models

var Cart []ItemDetails

type ItemDetails struct {
	ItemId          string `json:"itemId"`
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
	UserName        string `json:"username"`
}
