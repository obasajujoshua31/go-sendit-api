package database

import "sendit-api/api/models"

var users = []models.User{
	models.User{
		Name:     "John Doe",
		Email:    "johndoe@email.com",
		Password: "12345678",
	},
}

var parcels = []models.Parcel{
	models.Parcel{
		Name:           "My first parcel",
		Destination:    "Lekki Island",
		PickUpLocation: "Yaba Poly",
	},
}
