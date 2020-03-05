package utils

import (
	"database/sql"
	"log"
	"os"

	//mysql for used only for side effect

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

// ConnectDataBase used for database Connection
func ConnectDataBase(idleConnection int) *sql.DB {

	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUsername := os.Getenv("databaseUsername")
	databasePassword := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")

	// "kalana:Kalife@2019@/New_Test"
	log.Println("INFO: DataBase Connecting....")
	db, err := sql.Open("mysql", databaseUsername+":"+databasePassword+"@/"+databaseName)
	if err != nil {
		log.Fatal("ERROR: DataBase Connection has been Failed. ", err)
	}
	log.Println("INFO: DataBase Connected Successfully.")

	if idleConnection <= 0 {
		log.Fatal("ERROR: Number of Idle connection cannot be less than zero.")
	}

	db.SetMaxIdleConns(idleConnection)

	return db
}
