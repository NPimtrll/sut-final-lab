package entity

import (
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	Name       string `valid:"required~Need Name"`
	Age        string //`valid:"stringlength(0|2)~Age is not integer"`
	CustomerID string `valid:"matches(^[C]\\s[MU]\\d{8}$)~ComtomerID is not pattern"`
}
