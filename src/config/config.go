package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientID          string
	ClientSecret      string
	MessageDeployment string
	Region            string
	Environment       string
	Nickname          string
	Email             string
	FirstName         string
	LastName          string
}

func SetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientId := os.Getenv("clientId")
	clientSecret := os.Getenv("clientSecret")
	messageDeployment := os.Getenv("messageDeployment")
	region := os.Getenv("region")
	environment := os.Getenv("environment")
	nickname := os.Getenv("nickname")
	email := os.Getenv("email")
	firstName := os.Getenv("firstName")
	lastName := os.Getenv("lastName")

	config := Config{
		ClientID:          clientId,
		ClientSecret:      clientSecret,
		MessageDeployment: messageDeployment,
		Region:            region,
		Environment:       environment,
		Nickname:          nickname,
		Email:             email,
		FirstName:         firstName,
		LastName:          lastName,
	}

	return config
}
