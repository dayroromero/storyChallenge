package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dayroromero/storiChallenge/pkg/db"
	"github.com/dayroromero/storiChallenge/utils"
)

type Transaction struct {
	AccountID int
	Date      time.Time
	Amount    float64
	Type      string
}

func File_processor() {
	file, err := os.Open("transactions.csv")
	if err != nil {
		fmt.Println("Error opening CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		accountID, err := utils.Atoi(line[0])
		if err != nil {
			log.Println("Error triying to convert", line[0])
		}

		date, err := utils.ParseDate(line[1], "2/1/2006")
		if err != nil {
			log.Println("Error triying to convert", line[1])
		}

		amount, err := utils.ParseFloat64(line[2])
		if err != nil {
			log.Println("Error triying to convert", line[2])
		}

		var transactionType string
		if amount >= 0 {
			transactionType = "debit"
		} else {
			transactionType = "credit"
		}

		transaction := Transaction{
			AccountID: accountID,
			Date:      date,
			Amount:    amount,
			Type:      transactionType,
		}

		processTransaction(transaction)
	}
}

func processTransaction(transaction Transaction) {
	db := db.GetInstance().DB

	db.Create(&transaction)
}
