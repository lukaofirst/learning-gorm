package gorm_methods

import (
	"GoSandbox/entities"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) uuid.UUID {
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

	// Insert the customer into the database
	if err := db.Create(&customer).Error; err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}

	log.Printf("Create Customer - Id: %s", customer.Id)

	return id
}

func QueryAll(db *gorm.DB) {
	var customers []entities.Customer

	result := db.Find(&customers)

	if result.Error != nil {
		log.Fatalf("Failed to query customers: %v", result.Error)
	}

	for _, customer := range customers {
		log.Printf("GetAll - Customer Id: %s, Name: %s, Age: %d, CreatedAt: %s", customer.Id, customer.Name, customer.Age, customer.CreatedAt)
	}
}

func QueryById(db *gorm.DB, id uuid.UUID) {
	var customer entities.Customer
	// id, err := uuid.Parse("01945C42-0E76-7EE1-99B1-5F387E810DB3")

	// if err != nil {
	// 	log.Fatalf("Failed to parse Id: %v", err)
	// }

	result := db.First(&customer, "Id = ?", id)

	if result.Error != nil {
		log.Fatalf("Failed to query customer: %v", result.Error)
	}

	log.Printf("GetById - Customer Id: %s, Name: %s, Age: %d, CreatedAt: %s", customer.Id, customer.Name, customer.Age, customer.CreatedAt)
}

func Update(db *gorm.DB, id uuid.UUID) {
	var customer entities.Customer

	// id, err := uuid.Parse("01945C42-0E76-7EE1-99B1-5F387E810DB3")

	// if err != nil {
	// 	log.Fatalf("Failed to parse Id: %v", err)
	// }

	result := db.First(&customer, "Id = ?", id)

	if result.Error != nil {
		log.Fatalf("Failed to query customer: %v", result.Error)
	}

	log.Print(id)

	customer.Name = "Lorem Updated2"
	customer.Age = 26

	// if your Id field is type uuid, for safety reason, reassign the id here
	// that won't happen if type int
	customer.Id = id

	now := time.Now()
	customer.UpdatedAt = &now

	db.Save(&customer)

	log.Printf("Update - Customer Id: %s, Name: %s, Age: %d, CreatedAt: %s, UpdatedAt: %s", customer.Id, customer.Name, customer.Age, customer.CreatedAt, customer.UpdatedAt)
}

func DeleteById(db *gorm.DB, id uuid.UUID) {
	// id, err := uuid.Parse("01945C42-0E76-7EE1-99B1-5F387E810DB3")

	// if err != nil {
	// 	log.Fatalf("Failed to parse Id: %v", err)
	// }

	result := db.Delete(&entities.Customer{}, id)

	if result.Error != nil {
		log.Fatalf("Failed with delete customer: %v", result.Error)
	}

	log.Printf("Delete - Customer Id: %s", id)
}
