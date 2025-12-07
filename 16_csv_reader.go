package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Muhammad,Ferdynan,Ali,Syahbana\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	// perulangan yang tidak pernah berhenti, harus dikasih break
	for {
		record, err := reader.Read()
		if err == io.EOF { // EOF= End OF File
			break
		}

		fmt.Println(record)
	}
}
