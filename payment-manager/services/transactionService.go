package services

import (
	"fmt"
	"payment-manager/models"
	"time"
)

func ProcessTransaction(transaction *models.Transaction) error {
	fmt.Printf("Transaction processing started for: %+v\n", transaction)
	time.Sleep(30 * time.Second)
	transaction.Status = "processed"
	fmt.Printf("Transaction processed for: %+v\n", transaction)
	return nil
}
