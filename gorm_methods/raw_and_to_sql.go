package gorm_methods

import (
	"GoSandbox/entities"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) int {
	product := entities.Product{
		Name:        "product name",
		Description: "description",
	}

	if err := db.Create(&product).Error; err != nil {
		log.Fatalf("Failed to create Product: %v", err)
	}

	return int(product.ID)
}

func QueryProductUsingRaw(db *gorm.DB, id int) {
	var product entities.Product

	result := db.Raw("SELECT * FROM gosandbox.\"Products\" WHERE gosandbox.\"Products\".\"ID\" = ?", id).Scan(&product)

	if result.Error != nil {
		log.Fatalf("Failed to query Product: %v", result.Error)
	}

	jsonBytes, jsonErr := json.Marshal(product)

	if jsonErr != nil {
		log.Print(jsonErr)
	}

	jsonSerialized := string(jsonBytes)

	log.Print(string(jsonSerialized))
}

func GenerateSQL(db *gorm.DB, id int) {
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&entities.Product{}).First(&entities.Product{}, "gosandbox.\"Products\".\"ID\" = ?", id)

		//return tx.Create(&entities.Product{})
	})

	log.Println(sql)
}
