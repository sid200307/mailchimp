package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailworker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	//sending real images 
	// Gmail SMTP settings
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Your Gmail credentials
	from := "siddharthdoshi53@gmail.com"
	appPassword := "cfbt pmjk ztsl vovy"

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
		
		msg,err2:=executeTemplate(recipent)
		if err2!=nil{
			fmt.Printf("Worker :%d Error Pasring template for %s",id,recipent.Email)
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
