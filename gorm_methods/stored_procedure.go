package gorm_methods

import (
	"log"

	"gorm.io/gorm"
)

func CreateStoredProcedure(db *gorm.DB) {
	storedProcedureSQL := `
		CREATE OR REPLACE PROCEDURE gosandbox.update_product_description(IN prod_id INT, IN prod_description TEXT)
		LANGUAGE plpgsql
		AS $$
		BEGIN
			UPDATE gosandbox."Products"
			SET "Description" = prod_description
			WHERE gosandbox."Products"."ID" = prod_id;
		END;
		$$;
	`

	if err := db.Exec(storedProcedureSQL).Error; err != nil {
		log.Printf("Error creating stored procedure: %v", err)
	} else {
		log.Printf("Successfully created stored procedure")
	}
}

func ExecStoredProcedure(db *gorm.DB) {
	if err := db.Exec("CALL gosandbox.update_product_description(?, ?)", 1, "New description for product 1").Error; err != nil {
		log.Printf("Error executing stored procedure: %v", err)
	} else {
		log.Printf("Successfully executed stored procedure")
	}
}
