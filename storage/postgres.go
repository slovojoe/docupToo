package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/slovojoe/docupToo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

func ConnectDB() {
	//Loading environment variables
	//dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password,
		dbPort)

	//Opening connection to database
	Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Printf("An error occured connecting to db %e", err)
	} else {
		fmt.Println("Successfully connected to db")
	}

	//Close connection to db after main finishes
	//  defer db.Close()

	// err := Db.Migrator().DropTable(&models.Document{})
	// if err != nil {
	// 	fmt.Printf("An error occured dropping documents table %s", err)
	// }else{fmt.Println("Documents table was dropped")}

	//Make database migrations if they have not been made
	Db.AutoMigrate(&models.Document{})
	Db.AutoMigrate(&models.User{})

}
