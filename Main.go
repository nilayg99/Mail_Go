package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var (
	Mail_subject = "Test mail subject"
	Mail_body    = "Test mail body"
	app_pass     = ""
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	// Access the environment variables
	app_pass = os.Getenv("app_pass")
	To_id := os.Getenv("To_id")
	from_id := os.Getenv("from_id")
	mailSend(from_id, To_id, app_pass)

}
func mailSend(from_id, To_id, app_pass string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from_id)
	m.SetHeader("To", To_id)
	m.SetHeader("Subject", Mail_subject)
	m.SetBody("text/html", Mail_body)
	m.Attach("./goLogo.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, from_id, app_pass)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {
	envLoad()
	fmt.Printf("Mail send completed.")
}
