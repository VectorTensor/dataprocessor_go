package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ColumnInfo struct {
	ColumnName    string  `gorm:"column:column_name"`
	DataType      string  `gorm:"column:data_type"`
	IsNullable    string  `gorm:"column:is_nullable"`
	ColumnDefault *string `gorm:"column:column_default"`
}

type DbOrm struct {
	dbsession *gorm.DB
	db_url    string
}

func (s *DbOrm) Initial() {

	app_env := os.Getenv("APP_ENV")
	if app_env == "" {
		app_env = "DEV"
	}

	if app_env == "DEV" {

		err := godotenv.Load("models/.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

	}

	db_url := os.Getenv("DB_URL")

	s.db_url = db_url

	log.Printf("URL : %s", db_url)
	session, err := gorm.Open(postgres.Open(s.db_url), &gorm.Config{})
	s.dbsession = session
	if err != nil {

		log.Fatalf("Failed to connect to DB: %v", err)

	}

}

func (s *DbOrm) GetSchema(tableName string) *[]ColumnInfo {

	schema := "generated_data" // or other schema if different

	var columns []ColumnInfo
	query := `
		SELECT column_name, data_type, is_nullable, column_default
		FROM information_schema.columns
		WHERE table_schema = ? AND table_name = ?
		ORDER BY ordinal_position;
	`
	if err := s.dbsession.Raw(query, schema, tableName).Scan(&columns).Error; err != nil {
		log.Fatalf("Failed to fetch column info: %v", err)
	}

	for _, col := range columns {
		fmt.Printf("Column: %-20s Type: %-15s Nullable: %-8s Default: %v\n",
			col.ColumnName, col.DataType, col.IsNullable, col.ColumnDefault)
	}

	return &columns

}
