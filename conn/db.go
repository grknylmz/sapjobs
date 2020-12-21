package db

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//Init initialize the db connection
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DBNAME")
	port := os.Getenv("PG_PORT")

	dsn := "host=" + host + " " + "user=" + user + " " + "password=" + password + " " + "dbname=" + " " + dbName + " " + "port=" + port

	session, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		log.Printf("Connected to DB.")
		log.Info().Msg("This message appears when log level set to Debug or Info")
	}
	db = session
}

// GetDB function to return DB connection
func GetDB() *gorm.DB {
	return db
}
