package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	//"bufio"
	//"strings"
	"bufio"
	"strings"
)

func main() {
	csvfile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
	}

	totalans := 0
	scanner := bufio.NewScanner(os.Stdin)
	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		question, answer := record[0], record[1]
		fmt.Println(question)

		scanner.Scan()
		text := scanner.Text()

		if (answer == strings.TrimSpace(text)) {
			fmt.Println(answer)
			totalans += 1
		}
	}

	fmt.Println(totalans)
}
