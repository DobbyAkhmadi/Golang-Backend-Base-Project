package seeders

import (
	"backend/internal/app/region/models"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func initProvinceSeed() error {
	// Open the CSV file
	file, err := os.Open("./internal/storage/provinces.csv")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Read all contents of the CSV file
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		return err
	}

	// Convert the buffer to a string
	csvData := buf.String()

	// Define the BOM as a string
	bom := "\uFEFF"

	// Remove the BOM from the CSV data, if present
	csvData = strings.TrimPrefix(csvData, bom)

	// Create a CSV reader with a custom delimiter ';'
	newReader := csv.NewReader(bytes.NewBufferString(csvData))
	newReader.Comma = ';'

	// Read all records from the CSV file
	records, err := newReader.ReadAll()
	if err != nil {
		return err
	}
	// Create the Province repository
	//	provinceRepo := repository.NewRegionRepository(db)

	// Process the CSV records
	var provinces []models.Province
	for _, record := range records {
		// Convert the string to an integer
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}

		province := models.Province{
			ID:   id,
			Name: record[1],
		}
		provinces = append(provinces, province)
	}

	// Print the imported data
	for _, province := range provinces {
		fmt.Printf("ID: %d, Name: %s\n", province.ID, province.Name)
	}

	return nil

}
