package relationships

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	Id    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string
	Coach Coach
}

type Coach struct {
	Id     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name   string
	TeamId uuid.UUID
}

func CreateTeam(db *gorm.DB) uuid.UUID {
	coachId, coachIdErr := uuid.NewV7()
	teamId, teamIdErr := uuid.NewV7()

	if coachIdErr != nil || teamIdErr != nil {
		log.Fatalf("Failed to generate UUID")
	}

	team := Team{
		Id:   teamId,
		Name: "TeamOne",
	}

	if err := db.Create(&team).Error; err != nil {
		log.Fatalf("Failed to create team: %v", err)
	}

	log.Printf("Create Team - Id: %s", team.Id)

	coach := Coach{
		Id:     coachId,
		Name:   "Coach Name",
		TeamId: teamId,
	}

	if err := db.Create(&coach).Error; err != nil {
		log.Fatalf("Failed to create coach: %v", err)
	}

	log.Printf("Create Coach - Id %s with Team - Id %s", coach.Id, team.Id)

	return team.Id
}

func QueryTeamById(db *gorm.DB, teamId uuid.UUID) {
	var team Team

	err := db.Model(&Team{}).Preload("Coach").First(&team, "gosandbox.\"Teams\".\"Id\" = ?", teamId)

	if err.Error != nil {
		log.Fatalf("Failed to query customer: %v", err.Error)
	}

	jsonBytes, jsonErr := json.Marshal(team)

	if jsonErr != nil {
		log.Print(jsonErr)
	}

	jsonSerialized := string(jsonBytes)

	log.Print(string(jsonSerialized))
}
