package models

import "gorm.io/gorm"

type Denda struct {
	gorm.Model
	//ID       int    `json:"id" form:"id"`
	No_Peminjam string `json: "no_peminjam" form: "no_peminjam"`
	Telat       string `json: "telat" form: "telat"`
	Pembayaran  string `json: "pembayaran" form: "pembayaran"`
	Nominal     string `json: "nominal" form: "nominal"`
}
