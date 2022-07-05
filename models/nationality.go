package models

import "gorm.io/gorm"

type Nacionality struct {
	gorm.Model
	Country    string
	City       string
	adress     string
	postalCode string
}
