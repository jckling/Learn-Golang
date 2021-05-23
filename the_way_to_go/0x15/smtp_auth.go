package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth( // 认证
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"mail.example.com:25",             // 服务器
		auth,                              // 认证
		"sender@example.org",              // 发件人
		[]string{"recipient@example.net"}, // 收件人
		[]byte("This is the email body."), // 邮件内容
	)
	if err != nil {
		log.Fatal(err)
	}
}
