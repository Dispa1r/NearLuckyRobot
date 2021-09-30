package model

import "github.com/jinzhu/gorm"

type User struct {
	*gorm.Model
	TgAccount int `gorm:"type:int(64);unique_index;not null"`
	NearAccount string `gorm:"type:varchar(32);unique;"`
	Money float64`gorm:"type:float(5,2);default:'0'"`
	UserName string `gorm:"type:text;"`
}