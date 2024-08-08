package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Accounts []Account
}

type Account struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
	UserID  uint
	User    User
	History []PaymentHistory
}

type PaymentHistory struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	AccountID uint
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
	Account   Account
}
