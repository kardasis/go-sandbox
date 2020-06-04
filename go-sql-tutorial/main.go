package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
  godotenv.Load()
  db, err := sql.Open("postgres", os.Getenv("POSTGRES_PARAMETER_STRING"))

  if err!= nil {
    panic(err.Error())
  }
  defer db.Close()

  fmt.Println("successfully connected")

  insert, err := db.Query("INSERT INTO dummy VALUES('ARI')")
  if err != nil {
    panic(err.Error())
  }
  defer insert.Close()

  fmt.Println("successfully inserted into users")
}
