package seed

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"go-rest-api/internal/auth"
)

func SeedAdmin(db *sql.DB){
	email := "admin@jagad.com"
	password := "admin123"

	// Check if admin already exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM admins WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if admin exists: %v", err)
	}
	if exists {
		log.Println("Admin already exists, skipping seeding.")
		return
	}
	
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Insert the admin into the database
	_, err = db.Exec("INSERT INTO admins (email, password) VALUES ($1, $2)", email, hash)
	if err != nil {
		log.Fatalf("Failed to insert admin: %v", err)
	}
	log.Println("Admin seeded successfully.")
}