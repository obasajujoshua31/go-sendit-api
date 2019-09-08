package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBDRIVER   = ""
	DBUSER     = ""
	DBPASSWORD = ""
	DBNAME     = ""
	DBHOST     = ""
	SECRETKEY  = ""
	PORT       = 0
	DBURL      = ""
)

func Load() {
	var err error
	godotenv.Load()

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBUSER = os.Getenv("DB_USER")
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	DBNAME = os.Getenv("DB_NAME")
	DBHOST = os.Getenv("DB_HOST")
	SECRETKEY = os.Getenv("SECRET_KEY")
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		log.Println(err)
		PORT = 7000
	}

	DBURL = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DBHOST, DBUSER, DBNAME, DBPASSWORD)
}
