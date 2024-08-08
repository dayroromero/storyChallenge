package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dayroromero/storiChallenge/pkg/csv"
	"github.com/dayroromero/storiChallenge/pkg/notifications"
)

func main() {
	lambda.Start(handler)
}

// handler Captures the file upload to S3 and processes it
func handler(event events.S3Event) {
	log.Println("Init Event")

	for _, record := range event.Records {
		s3entity := record.S3
		bucket := s3entity.Bucket.Name
		key := s3entity.Object.Key

		csv.File_processor(bucket, key)

		log.Printf("File %v processed succesfully", key)
	}

	// SendSummary for unique account
	notifications.SendSummary(1)
}
