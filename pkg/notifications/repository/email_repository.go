package emailrepository

import (
	"log"

	"github.com/dayroromero/storiChallenge/pkg/db"
	"github.com/dayroromero/storiChallenge/pkg/models"
)

func GetUser(userID int) (models.User, error) {
	var user models.User

	db := db.GetInstance().DB
	if err := db.First(&user).Error; err != nil {
		log.Printf("Error getting user from database: %v", err)
		return user, err
	}

	return user, nil
}

func GetSummary(userId int) models.EmailData {
	client, err := GetUser(userId)
	if err != nil {
		log.Println("Error getting User")
	}

	clientName := client.UserName

	totalBalance := getTotalBalanceByAccount(userId)

	transactionsJuly := getTransactionCountByMonth(7, userId)

	transactionsAugust := getTransactionCountByMonth(8, userId)

	averageDebitAmmount := getAverageDebitsCreditsByAccount("debit", userId)

	averageCreditAmmount := getAverageDebitsCreditsByAccount("credit", userId)

	data := models.EmailData{
		ClientName:           clientName,
		TotalBalance:         totalBalance,
		TransactionsInJuly:   transactionsJuly,
		TransactionsInAugust: transactionsAugust,
		AverageDebitAmount:   averageDebitAmmount,
		AverageCreditAmount:  averageCreditAmmount,
	}

	return data
}

func getTransactionCountByMonth(month int, accountID int) int64 {
	db := db.GetInstance().DB

	var transactionCount int64

	if err := db.Model(&models.Transaction{}).
		Where("account_id = ? AND EXTRACT(MONTH FROM date) = ?", accountID, month).
		Count(&transactionCount).Error; err != nil {
		log.Printf("Error to count transactiosn: %v", err)
		return 0
	}

	return transactionCount
}

func getTotalBalanceByAccount(accountID int) float64 {
	db := db.GetInstance().DB

	var totalBalance float64
	if err := db.Table("transactions").Where("account_id = ?", 1).Select("SUM(amount)").Row().Scan(&totalBalance); err != nil {
		log.Printf("Error calculating the sum of transaction amounts: %v", err)
		return 0
	}

	return totalBalance
}

func getAverageDebitsCreditsByAccount(transactionType string, accountID int) float64 {
	db := db.GetInstance().DB

	var totalAmount float64
	if err := db.Model(&models.Transaction{}).
		Where("account_id = ? AND type = ?", accountID, transactionType).
		Select("AVG(amount)").Row().Scan(&totalAmount); err != nil {
		log.Printf("Error gettin debit/credit by account: %v", err)
		return 0
	}

	return totalAmount
}
