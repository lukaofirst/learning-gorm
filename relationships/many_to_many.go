package relationships

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamMtm struct {
	Id      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name    string
	Matches []Match `gorm:"many2many:TeamMatches;joinForeignKey:TeamMtmId;joinReferences:MatchId;constraint:OnDelete:CASCADE;"`
}

type TeamMatch struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;not null;"`
	TeamMtmId uuid.UUID `gorm:"type:uuid;foreignKey:TeamMtmId;references:Id;constraint:OnDelete:CASCADE;"`
	MatchId   uuid.UUID `gorm:"type:uuid;foreignKey:MatchId;references:Id;constraint:OnDelete:CASCADE;"`
}

type Match struct {
	Id    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string
	Teams []TeamMtm `gorm:"many2many:TeamMatches;joinForeignKey:MatchId;joinReferences:TeamMtmId;constraint:OnDelete:CASCADE;"`
}

func CreateTeamMTMAndMatches(db *gorm.DB) {
	var teamMtms []TeamMtm
	for i := 0; i < 3; i++ {
		teamMtmId, teamMtmIdErr := uuid.NewV7()

		if teamMtmIdErr != nil {
			log.Fatalf("Failed to generate UUID")
		}

		teamMtm := TeamMtm{
			Id:   teamMtmId,
			Name: fmt.Sprintf("TeamMtm %d", i),
		}

		teamMtms = append(teamMtms, teamMtm)
	}

	if err := db.Create(&teamMtms).Error; err != nil {
		log.Fatalf("Failed to create TeamMtms: %v", err)
	}

	for _, match := range teamMtms {
		log.Printf("Create TeamMtms - Id: %s", match.Id)
	}

	var matches []Match
	for i := 0; i < 3; i++ {
		matchId, matchIdErr := uuid.NewV7()

		if matchIdErr != nil {
			log.Fatalf("Failed to generate UUID")
		}

		match := Match{
			Id:   matchId,
			Name: fmt.Sprintf("Match %d", i),
		}

		matches = append(matches, match)
	}

	if err := db.Create(&matches).Error; err != nil {
		log.Fatalf("Failed to create matches: %v", err)
	}

	for _, match := range matches {
		log.Printf("Create Match - Id: %s", match.Id)
	}
}

func CreateTeamMatch(db *gorm.DB) {
	var teamMtms []TeamMtm
	var matches []Match

	teamMtmsResult := db.Limit(1).Model(&TeamMtm{}).Find(&teamMtms)
	matchesResult := db.Limit(1).Model(&Match{}).Find(&matches)

	if teamMtmsResult.Error != nil || matchesResult.Error != nil {
		log.Fatalf("Failed to query TeamMtms and Matches %v, %v", teamMtmsResult.Error, matchesResult.Error)
	}

	var teamMatches []TeamMatch
	for _, teamMtm := range teamMtms {
		for _, match := range matches {
			teamMatchId, teamMatchIdErr := uuid.NewV7()

			if teamMatchIdErr != nil {
				log.Fatalf("Failed to generate UUID")
			}

			teamMatch := TeamMatch{
				Id:        teamMatchId,
				TeamMtmId: teamMtm.Id,
				MatchId:   match.Id,
			}

			teamMatches = append(teamMatches, teamMatch)
		}
	}

	if err := db.Create(&teamMatches).Error; err != nil {
		log.Fatalf("Failed to create teamMatches: %v", err)
	}

	for _, teamMatch := range teamMatches {
		log.Printf("Create TeamMatch - Id: %s", teamMatch.Id)
	}
}

func QueryTeamMatches(db *gorm.DB) {
	var teamMatches []TeamMatch

	err := db.Model(&TeamMatch{}).Find(&teamMatches)

	if err.Error != nil {
		log.Fatalf("Failed to query teamMatches: %v", err.Error)
	}

	jsonBytes, jsonErr := json.Marshal(teamMatches)

	if jsonErr != nil {
		log.Print(jsonErr)
	}

	jsonSerialized := string(jsonBytes)

	log.Print(string(jsonSerialized))
}
