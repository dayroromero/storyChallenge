package csv

import (
	"encoding/csv"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dayroromero/storiChallenge/pkg/db"
	"github.com/dayroromero/storiChallenge/utils"
)

type Transaction struct {
	AccountID int
	Date      time.Time
	Amount    float64
	Type      string
}

func File_processor(bucket string, key string) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		return
	}
	defer result.Body.Close()

	reader := csv.NewReader(result.Body)

	lines, err := reader.ReadAll()
	if err != nil {
		return
	}

	for i, line := range lines {
		if i == 0 {
			continue
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
			AccountID: 1,
			Date:      date,
			Amount:    amount,
			Type:      transactionType,
		}
		log.Println(transaction)
		saveTransaction(transaction)
	}
}

func saveTransaction(transaction Transaction) {
	log.Println("Saving Transactions into database")
	db := db.GetInstance().DB

	db.Create(&transaction)
}
