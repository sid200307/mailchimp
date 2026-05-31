package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"
)

func emailworker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	//sending real images
	// Gmail SMTP settings
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	//Gamil Crendentials
	from := os.Getenv("EMAIL_FROM")
	appPassword := os.Getenv("EMAIL_APP_PASSWORD")

	auth := smtp.PlainAuth(
		"",
		from,
		appPassword,
		smtpHost,
	)

	//Testing Service Mailpit
	// smtpHost := "localhost"
	// smtpPort := "1025"
	// formatedmsg := fmt.Sprintf(
	// 		"To:%s\r\n"+"Subject: Test Email \r\n\r\n%s\r\n",
	// 		recipent.Email,
	// 		"Just testing out Testing email Campagin",
	// 	)

	// msg := []byte(formatedmsg)

	for recipent := range ch {

		msg, err2 := executeTemplate(recipent)
		if err2 != nil {
			fmt.Printf("Worker :%d Error Pasring template for %s", id, recipent.Email)
			//add to dlq
			continue
		}

		fmt.Printf("Worker-%d: Sending email to %s\n", id, recipent.Email)
		//using auth sending real mails
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipent.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker-%d: Sent email to %s\n", id, recipent.Email)
	}

}
