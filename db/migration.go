// Package db
package db

import (
	"flight-booking-server/entities"
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	// db.Exec("CREATE TYPE ticket_status AS ENUM ('approved', 'cancelled')")
	// db.Migrator().DropTable(&entities.Ticket{})
	err := db.AutoMigrate(
		&entities.Ticket{},
		&entities.Seat{},
		&entities.Plane{},
		&entities.Location{},
		&entities.Flight{},
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
		panic(err)
	}

	log.Println("Database migration completed")
}
