package relationships

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamOtm struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	LeagueId uuid.UUID
}

type League struct {
	Id    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string
	Teams []TeamOtm
}

func CreateLeague(db *gorm.DB) uuid.UUID {
	leagueId, leagueIdErr := uuid.NewV7()

	if leagueIdErr != nil {
		log.Fatalf("Failed to generate UUID")
	}

	league := League{
		Id:   leagueId,
		Name: "LeagueOne",
	}

	if err := db.Create(&league).Error; err != nil {
		log.Fatalf("Failed to create league: %v", err)
	}

	log.Printf("Create League - Id: %s", league.Id)

	var teams []TeamOtm
	for i := 0; i < 3; i++ {
		teamId, teamIdErr := uuid.NewV7()

		if teamIdErr != nil {
			log.Fatalf("Failed to generate UUID")
		}

		team := TeamOtm{
			Id:       teamId,
			Name:     "Team Name",
			LeagueId: leagueId,
		}

		teams = append(teams, team)
	}

	if err := db.Create(&teams).Error; err != nil {
		log.Fatalf("Failed to create teams: %v", err)
	}

	for _, team := range teams {
		log.Printf("Create Team - Id: %s", team.Id)
	}

	return league.Id
}

func QueryLeagueById(db *gorm.DB, leagueId uuid.UUID) {
	var league League

	err := db.Model(&League{}).Preload("Teams").First(&league, "gosandbox.\"Leagues\".\"Id\" = ?", leagueId)

	if err.Error != nil {
		log.Fatalf("Failed to query customer: %v", err.Error)
	}

	jsonBytes, jsonErr := json.Marshal(league)

	if jsonErr != nil {
		log.Print(jsonErr)
	}

	jsonSerialized := string(jsonBytes)

	log.Print(string(jsonSerialized))
}
