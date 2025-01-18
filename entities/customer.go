package entities

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt *time.Time `gorm:"autoUpdateTime:false"`
}

// If you want to set a specific table's name
// func (Customer) TableName() string {
// 	return "Customers" // This sets the table name to 'Customers'
// }
