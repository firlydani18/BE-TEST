package services

import (
	"account-manager/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return db.Create(user).Error
}

func AuthenticateUser(db *gorm.DB, user *models.User) error {
	var foundUser models.User
	if err := db.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
}
