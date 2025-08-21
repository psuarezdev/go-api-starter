package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/*import (
	"os"
	"github.com/joho/godotenv"
)*/

const API_PREFIX = "/api/v1"

func GetConnectionString() string {
	godotenv.Load(".env")

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatal("Error parsing DB_PORT:", err)
		panic(err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
}
