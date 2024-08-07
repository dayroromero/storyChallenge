package notifications

type EmailData struct {
	ClientName           string
	TotalBalance         float64
	TransactionsInJuly   int
	TransactionsInAugust int
	AverageDebitAmount   float64
	AverageCreditAmount  float64
}
