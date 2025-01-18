package gorm_methods

import (
	"GoSandbox/entities"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCustomer(db *gorm.DB, name string, age int, createdAt time.Time) {
	id, err := uuid.NewV7()

	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}

	customer := entities.Customer{
		Id:        id,
		Name:      name,
		Age:       age,
		CreatedAt: createdAt,
		UpdatedAt: nil,
	}

	// Insert the customer into the database
	if err := db.Create(&customer).Error; err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}
}

func GetCustomers(db *gorm.DB, page, pageSize int, filter, searchTerm, sortBy, sortOrder string) {
	query := db

	if filter != "" {
		query = query.Where("gosandbox.\"Customers\".\"Name\" LIKE ?", "%"+filter+"%")
	}

	if searchTerm != "" {
		query = query.Where("gosandbox.\"Customers\".\"Name\" LIKE ?", "%"+searchTerm+"%")
	}

	if sortBy != "" {
		query = query.Order(sortBy + " " + sortOrder)
	}

	var customers []entities.Customer

	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&customers).Error; err != nil {
		log.Fatalf("Failed to query customers: %v", err)
	}

	jsonBytes, jsonErr := json.Marshal(customers)

	if jsonErr != nil {
		log.Print(jsonErr)
	}

	jsonSerialized := string(jsonBytes)

	log.Print(string(jsonSerialized))
}
