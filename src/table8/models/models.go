package models

import (
	"time"
)

/**
* Common data models for several API requests/responses
**/

////////// FOR GORM TESTING ONLY

type Region struct {
	Uid       	string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	Title		string
}

type RestaurantTag struct {
	Uid       	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	Title		string
}

type Restaurant struct {
	Uid       	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	Title		string
	Regions		[]string
	Tags		[]string
	Details		string
}
