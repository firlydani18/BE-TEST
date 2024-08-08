package models

type Transaction struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Timestamp string  `json:"timestamp"`
	ToAddress string  `json:"to_address"`
	Status    string  `json:"status"`
}
