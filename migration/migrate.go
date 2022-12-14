package migration

import (
	cartModel "capstone/happyApp/features/cart/data"
	eventModel "capstone/happyApp/features/event/data"

	// userModel "capstone/happyApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(cartModel.User{})
	db.AutoMigrate(cartModel.Community{})
	db.AutoMigrate(cartModel.JoinCommunity{})
	db.AutoMigrate(eventModel.Event{})
	db.AutoMigrate(eventModel.JoinEvent{})
	db.AutoMigrate(cartModel.Feed{})
	db.AutoMigrate(cartModel.Comment{})
	db.AutoMigrate(cartModel.Product{})
	db.AutoMigrate(cartModel.Cart{})
	db.AutoMigrate(cartModel.Transaction{})
	db.AutoMigrate(cartModel.TransactionCart{})
	db.AutoMigrate(cartModel.Payment{})
	//
}

// do nothing
// do nothing
