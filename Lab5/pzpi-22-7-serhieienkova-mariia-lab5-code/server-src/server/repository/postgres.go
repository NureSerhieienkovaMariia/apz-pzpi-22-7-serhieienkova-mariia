package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	relativesTable               = "public.\"relatives\""
	doctorsTable                 = "public.\"doctors\""
	adminsTable                  = "public.\"admins\""
	patientsTable                = "public.\"patients\""
	medicinesTable               = "public.\"medicines\""
	diagnosesTable               = "public.\"diagnoses\""
	patientMedicineTable         = "public.\"patients_medicines\""
	userPatientsTable            = "public.\"users_patients\""
	devicesTable                 = "public.\"device\""
	indicatorsNotificationsTable = "public.\"indicators_notifications\""
	indicatorsStampsTable        = "public.\"indicators_stamps\""
	patientNotesTable            = "public.\"patient_notes\""
	patientRelativeTable         = "public.\"patients_relatives\""
	prescriptionsTable           = "public.\"prescription\""
	treatmentPlansTable          = "public.\"treatment_plans\""
	visitsTable                  = "public.\"visits\""
	patientsRelativesTable       = "public.\"patients_relatives\""
	patientsDiagnosesTable       = "public.\"patients_diagnoses\""
	patientNotificationTable     = "public.\"patient_notification\""
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type TableData struct {
	TableName string      `json:"table_name"`
	Data      interface{} `json:"data"`
}

func ExportAllDataToJSON(db *sqlx.DB, outputPath string) error {
	tables := []string{
		relativesTable,
		doctorsTable,
		adminsTable,
		patientsTable,
		medicinesTable,
		diagnosesTable,
		patientMedicineTable,
		userPatientsTable,
		devicesTable,
		indicatorsNotificationsTable,
		indicatorsStampsTable,
		patientNotesTable,
		patientRelativeTable,
		prescriptionsTable,
		treatmentPlansTable,
		visitsTable,
		patientsRelativesTable,
		patientsDiagnosesTable,
		patientNotificationTable,
	}

	var allData []TableData

	for _, table := range tables {
		var rows []map[string]interface{}
		query := fmt.Sprintf("SELECT * FROM %s", table)
		if err := db.Select(&rows, query); err != nil {
			return fmt.Errorf("failed to fetch data from table %s: %w", table, err)
		}
		allData = append(allData, TableData{
			TableName: table,
			Data:      rows,
		})
	}

	jsonData, err := json.MarshalIndent(allData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write JSON data to file: %w", err)
	}

	return nil
}

func ImportDataFromJSON(db *sqlx.DB, inputPath string) error {
	// Read the JSON file
	jsonData, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse the JSON data
	var allData []TableData
	if err := json.Unmarshal(jsonData, &allData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	// Insert data into each table
	for _, tableData := range allData {
		for _, row := range tableData.Data.([]interface{}) {
			rowMap := row.(map[string]interface{})

			// Build the INSERT query dynamically
			columns := []string{}
			values := []interface{}{}
			placeholders := []string{}
			for col, val := range rowMap {
				columns = append(columns, col)
				values = append(values, val)
				placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
			}

			query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
				tableData.TableName,
				strings.Join(columns, ", "),
				strings.Join(placeholders, ", "))

			// Execute the query
			if _, err := db.Exec(query, values...); err != nil {
				return fmt.Errorf("failed to insert data into table %s: %w", tableData.TableName, err)
			}
		}
	}

	return nil
}
