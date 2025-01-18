package gorm_methods

import (
	"GoSandbox/entities"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RunTransaction(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		id, err := uuid.NewV7()

		if err != nil {
			log.Fatalf("Failed to generate UUID: %v", err)
		}

		customer := entities.Customer{
			Id:        id,
			Name:      "Lorem",
			Age:       18,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}

		if err := tx.Create(&customer).Error; err != nil {
			log.Fatalf("Failed to create customer: %v", err)
			return err
		}

		log.Printf("Create Customer - Id: %s", customer.Id)

		if err := tx.First(&customer).Error; err != nil {
			log.Fatalf("Failed to query customer: %v", err)
			return err
		}

		customer.Name = "Lorem Updated2"
		customer.Age = 26

		// if your Id field is type uuid, for safety reason, reassign the id here
		// that won't happen if type int
		customer.Id = id

		now := time.Now()
		customer.UpdatedAt = &now

		if err := tx.Save(&customer).Error; err != nil {
			log.Fatalf("Failed to save customer: %v", err)
			return err
		}

		log.Printf("Update - Customer Id: %s, Name: %s, Age: %d, CreatedAt: %s, UpdatedAt: %s", customer.Id, customer.Name, customer.Age, customer.CreatedAt, customer.UpdatedAt)

		return nil
	})
}
