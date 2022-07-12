package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {

	err := godotenv.Load(".env")
	require.NoError(t, err)

	User := os.Getenv("DB_USER")
	Pass := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")

	require.NotEmpty(t, User)
	require.NotEmpty(t, Pass)
	require.NotEmpty(t, DbName)
	require.NotEmpty(t, Host)
	require.NotEmpty(t, Port)

	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Mexico_City",
		Host, User, Pass, DbName, Port)

	_, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	require.NoError(t, err)
}
