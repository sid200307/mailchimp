package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(filepath string, ch chan Recipient) error {
	f, error := os.Open(filepath)
	if error != nil {
		return error

	}
	// after ending of loadreciepent function defer will execure close nno matter eroor or what
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return error

	}
	for _, rec := range records[1:] {

		//send ->consumer ->channels
		ch <- Recipient{
			Name:  rec[0],
			Email: rec[1],
		}

	}

	return nil
}
