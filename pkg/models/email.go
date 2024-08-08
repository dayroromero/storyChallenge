package models

type EmailData struct {
	ClientName           string
	TotalBalance         float64
	TransactionsInJuly   int64
	TransactionsInAugust int64
	AverageDebitAmount   float64
	AverageCreditAmount  float64
}
