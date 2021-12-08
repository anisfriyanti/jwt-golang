package models

type Transaction struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	User_Id uint `json:"user_id"`
	Product_Id uint `json:"product_id"`
	Amount uint `json:"amount"`
	Status string `json:"status"`


}