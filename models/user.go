package models

import ()

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age" gorm:default:0`
	City     string `json:"city"`
}
