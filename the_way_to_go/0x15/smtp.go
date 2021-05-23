package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func main() {
	// Connect to the remote SMTP server.
	client, err := smtp.Dial("mail.example.com:25") // SMTP 服务器客户端
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient.
	client.Mail("sender@example.org")    // 发件人
	client.Rcpt("recipient@example.net") // 收件人

	// Send the email body.
	wc, err := client.Data() // 用于写入数据
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
