package main

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

const (
	filepath = `D:\GO\mailchimp\users_1000.csv`
)

func main() {

	fmt.Println("Welcome to email dispatcher!")

	recipentchan := make(chan Recipient)
	//waitgroup
	var wg sync.WaitGroup
	//worker pool
	workercount := 10

	go func() {
		defer close(recipentchan) //always close channel else deadlock will happen
		err := loadRecipient(filepath, recipentchan)
		if err != nil {
			return
		}
	}()

	for i := 1; i <= workercount; i++ {
		wg.Add(1)
		go emailworker(i, recipentchan, &wg)
	}

	wg.Wait()

}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	errr := t.Execute(&tpl, r)
	if errr != nil {

		return "", errr

	}

	return tpl.String(), nil
}
