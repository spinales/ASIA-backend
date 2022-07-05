package util

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI               string
	AppVersion          string
	DBVersion           uint
	ServerAddress       string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessTokenDuration := os.Getenv("Access_Token_Duration")
	num, err := strconv.Atoi(accessTokenDuration)
	if err != nil {
		log.Panicln("Invalid time unit in .env file, Access Token Duration")
		return Config{}
	}

	return Config{
		DBURI:               os.Getenv("DATABASE_URL"),
		AppVersion:          "v0.1",
		DBVersion:           1, // current db version
		ServerAddress:       os.Getenv("Server_Address"),
		TokenSymmetricKey:   "ZnQSxSQRpZXt2ATNt5ktZ6bamCY6JHEn", // pending
		AccessTokenDuration: time.Hour * time.Duration(num),
	}
}
