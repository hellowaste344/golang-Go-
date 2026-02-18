package main

import (
	"fmt"
	"net/smtp"
	"os"
	"log"
    "github.com/joho/godotenv"
)


func emailSender(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("EMAIL")
	secretKey := os.Getenv("EMAIL_PASS")
	fmt.Println(s3Bucket, secretKey)
	toList := []string{"example@gmail.com"}

	host := "smtp.gmail.com"
	port := "587"
	msg := "Subject: Test Email\r\n\r\nYour message here..."

	auth := smtp.PlainAuth("", s3Bucket, secretKey, host)

	err = smtp.SendMail(host+":"+port, auth, s3Bucket, toList, []byte(msg))
	
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	fmt.Println("Successfully sent an email")
}