package models

import (
	"os"
	"fmt"
  "log"

  _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

  e := godotenv.Load()

  if e != nil {
    log.Println("ENV: ", e)
  }

  username := os.Getenv("POSTGRES_USER")
  password := os.Getenv("POSTGRES_PASSWORD")
  dbName := os.Getenv("POSTGRES_DB")
  dbHost := os.Getenv("db_host")

  //Build connection string
  dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
  log.Println(dbUri)

  conn, err := gorm.Open("postgres", dbUri)

  if err != nil {
    log.Print("Error:", err, "\n")
  }

  db = conn
  db.Debug().AutoMigrate(&Invoice{}, &Charge{}) // DB migration # TODO
}

func GetDB() *gorm.DB {
  return db
}
