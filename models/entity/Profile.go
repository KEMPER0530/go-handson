package entity

type Profile struct {
	Id       int    `gorm:"primary_key;not null"     json:"id"`
	Heading  string `gorm:"type:varchar(100);null"   json:"heading"`
	Lastdate string `gorm:"type:null"                json:"lastdate"`
	History  string `gorm:"type:varchar(100);null"   json:"history"`
}
