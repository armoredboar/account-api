package main

import (
	"log"
	"os"

	"github.com/armoredboar/account-api/server"
)

func main() {

	if _, exists := os.LookupEnv("AB_EMAIL_SMTP"); exists == false {
		log.Fatal("AB_EMAIL_SMTP environment variable is not set.")
	}

	if _, exists := os.LookupEnv("AB_EMAIL_PORT"); exists == false {
		log.Fatal("AB_EMAIL_PORT environment variable is not set.")
	}

	if _, exists := os.LookupEnv("AB_EMAIL_NOREPLY"); exists == false {
		log.Fatal("AB_EMAIL_NOREPLY environment variable is not set.")
	}

	if _, exists := os.LookupEnv("AB_EMAIL_NOREPLY_PASSWORD"); exists == false {
		log.Fatal("AB_EMAIL_NOREPLY_PASSWORD environment variable is not set.")
	}

	accountApi.Server.Run()
}
