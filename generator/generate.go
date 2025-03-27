package main

// Generates a 2 column csv to be used with 3dp.yt qr redirector
// Columns: URL, Label
// Example: https://3dp.yt/qr/12345678,12345678
// The ID will be All uppercase, alpha numeric, and 8 characters long
// The URL will be the full URL to the QR code
// The Label will be the ID

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Generate a random alphanumeric string of length 8
func generateID() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 8; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}

// Create the CSV file with URL and Label columns
func createCSV(filename string, numEntries int) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"URL", "Label"})

	// Write entries
	for i := 0; i < numEntries; i++ {
		id := generateID()
		url := fmt.Sprintf("https://3dp.yt/fq/%s", id) //Change to /custom/ if you want to use custom qr codes
		label := id

		err := writer.Write([]string{url, label})
		if err != nil {
			return fmt.Errorf("could not write to CSV: %w", err)
		}
	}

	return nil
}

func main() {
	// Number of entries you want to generate
	numEntries := 100 // You can change this to generate more or fewer rows
	filename := "qr_codes.csv"

	// Create the CSV file
	err := createCSV(filename, numEntries)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("CSV file '%s' has been created successfully!\n", filename)
}
