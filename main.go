package main

import (
	"GoSandbox/entities"
	"GoSandbox/gorm_methods"
	relationships "GoSandbox/relationships"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// Initial connection string for the default "master" database
	db := initDB()

	// Basic CRUD
	// id := gorm_methods.Create(db)
	// gorm_methods.QueryAll(db)
	// gorm_methods.QueryById(db, id)
	// gorm_methods.Update(db, id)
	// gorm_methods.DeleteById(db, id)

	// Using Transactions
	// gorm_methods.RunTransaction(db)

	// Relationships - One to One
	//teamId := relationships.CreateTeam(db)
	//relationships.QueryTeamById(db, teamId)

	// Relationships - One to Many
	//leagueId := relationships.CreateLeague(db)
	//relationships.QueryLeagueById(db, leagueId)

	// Relationships - Many to Many
	// relationships.CreateTeamMTMAndMatches(db)
	// relationships.CreateTeamMatch(db)
	// relationships.QueryTeamMatches(db)

	// Relationships - Raw SQL and ToSQL
	// id := gorm_methods.CreateProduct(db)
	// gorm_methods.QueryProductUsingRaw(db, id)
	// gorm_methods.GenerateSQL(db, id)

	// Creating Views
	// gorm_methods.CreateView(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.QueryUsingView(db)

	// Creating Stored Procedures
	// gorm_methods.CreateStoredProcedure(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.CreateProduct(db)
	// gorm_methods.ExecStoredProcedure(db)

	// Paging, Filtering, Searching and Ordering
	gorm_methods.CreateCustomer(db, "Lorem", 18, time.Now().Add(24*time.Hour))
	gorm_methods.CreateCustomer(db, "John", 20, time.Now().Add(-24*time.Hour))
	gorm_methods.CreateCustomer(db, "Ipsum", 25, time.Now().Add(48*time.Hour))
	gorm_methods.CreateCustomer(db, "Jane", 32, time.Now().Add(96*time.Hour))
	gorm_methods.CreateCustomer(db, "Dolor", 40, time.Now().Add(-128*time.Hour))
	gorm_methods.GetCustomers(db, 1, 10, "", "", "gosandbox.\"Customers\".\"Name\"", "ASC")
}

func initDB() *gorm.DB {
	dsn := "connection_string_here"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// this will turn "Customer" to "Customers"
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   true,
			TablePrefix:   "gosandbox.", // to use the "gosandbox" schema instead of default public
		},
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS gosandbox;").Error; err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	if err := db.AutoMigrate(
		&entities.Customer{},
		&relationships.Team{},
		&relationships.Coach{},
		&relationships.TeamOtm{},
		&relationships.League{},
		&relationships.TeamMtm{},
		&relationships.Match{},
		&relationships.TeamMatch{},
		&entities.Product{},
	); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	return db
}
