package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"

	gomail "gopkg.in/mail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	password := os.Getenv("MAIL_PASS")
	body, err := ioutil.ReadFile("email.html")
	if err != nil {
		fmt.Println(err)
	}

	abc := gomail.NewMessage()

	abc.SetHeader("From", "nihaltmdev@gmail.com")
	abc.SetHeader("To", "nihaltm2002@gmail.com")

	abc.SetHeader("Subject", "Test")
	abc.SetBody("text/html", string(body))
	// fmt.Println(password)

	a := gomail.NewDialer("smtp.gmail.com", 587, "nihaltmdev@gmail.com", password)

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
