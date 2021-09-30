package model

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	*gorm.Model
	Txn string `gorm:"type:varchar(256);unique"`
}
