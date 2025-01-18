package gorm_methods

import (
	"GoSandbox/entities"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

func CreateView(db *gorm.DB) {
	query := db.Model(&entities.Product{}).Where("gosandbox.\"Products\".\"ID\" > ?", 3)

	db.Migrator().CreateView("ProductView", gorm.ViewOption{Query: query, Replace: true})

	//db.Migrator().DropView("ProductView")
}

func QueryUsingView(db *gorm.DB) {
	var products []entities.Product

	if err := db.Table("ProductView").Find(&products).Error; err != nil {
		log.Printf("Error querying the view: %v", err)
	}

	jsonBytes, jsonErr := json.Marshal(products)

	if jsonErr != nil {
		log.Print(jsonErr)
	}

	jsonSerialized := string(jsonBytes)

	log.Print(string(jsonSerialized))
}
