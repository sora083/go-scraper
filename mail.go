package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendMail(body string) {
	fmt.Println("START mailSend: ")

	from := "test@gmail.com"
	to := "test@gmail.com"
	host := "localhost"

	// TODO メールの文字化け対応を入れる
	msg := []byte("" +
		"From: 送信した人 <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 件名 subject です\r\n\r\n" +
		body +
		"\r\n" +
		"")

	// func SendMail without auth
	if err := smtp.SendMail(host+":25", nil, from, []string{to}, msg); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("END mailSend: ")
}
