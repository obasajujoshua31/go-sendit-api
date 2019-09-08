package database

import (
	"log"
	"sendit-api/api/console"
	"sendit-api/api/models"
)

func MigrateAndSeed() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().DropTableIfExists(&models.User{}, &models.Parcel{}).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Parcel{}).Error

	if err != nil {
		log.Fatal(err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			log.Fatal(err)
		}

		err = db.Debug().Model(&models.Parcel{}).AddForeignKey("owner_id", "users(id)", "cascade", "cascade").Error

		if err != nil {
			log.Fatal(err)
		}

		parcels[i].OwnerID = users[i].ID

		err = db.Debug().Model(&models.Parcel{}).Create(&parcels[i]).Error

		if err != nil {
			log.Fatal(err)
		}

		err = db.Debug().Model(&models.Parcel{}).Where("owner_id = ?", parcels[i].OwnerID).Take(&models.Parcel{}).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(parcels[i])
	}

}
