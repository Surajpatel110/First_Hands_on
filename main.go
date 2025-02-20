package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to list all records in the CSV file
func listCSV(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file closes after function execution

	reader := csv.NewReader(file)    // Creating a new CSV reader
	records, err := reader.ReadAll() // Reading all records from the CSV file
	if err != nil {
		fmt.Printf("Error reading the CSV file: %v\n", err)
		return
	}

	// Iterate over each record and print it with an index
	for i, record := range records {
		fmt.Printf("Row %d: %v\n", i+1, record)
	}
}

// Function to add a new record to the CSV file with user input
func addRecord(filepath string) {
	// Opening the file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	reader := bufio.NewReader(os.Stdin) // Created a buffered reader to read user input

	fmt.Print("Enter SiteID: ")
	siteID, _ := reader.ReadString('\n')
	fmt.Print("Enter FixletID: ")
	fixletID, _ := reader.ReadString('\n')
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter Criticality: ")
	criticality, _ := reader.ReadString('\n')
	fmt.Print("Enter RelevantComputerCount: ")
	relevantCount, _ := reader.ReadString('\n')

	// Trim spaces and newlines from user input
	newRecord := []string{
		strings.TrimSpace(siteID),
		strings.TrimSpace(fixletID),
		strings.TrimSpace(name),
		strings.TrimSpace(criticality),
		strings.TrimSpace(relevantCount),
	}

	if err := writer.Write(newRecord); err != nil {
		log.Fatal("Error writing to file: ", err)
	}

	writer.Flush() // Ensure all buffered data is written to the file
	fmt.Println("Record added successfully!")
}

// Function to delete a record by FixletID
func deleteRecord(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll() // Read all records into memory
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	if len(records) == 0 {
		fmt.Println("No records to delete.")
		return
	}

	var fixletIDToDelete string
	fmt.Print("Enter FixletID to delete: ")
	fmt.Scanln(&fixletIDToDelete)

	newRecords := [][]string{} // Store the filtered records
	deleted := false           // Flag to track if deletion occurred

	// Iterate through records, keeping all except the matching FixletID
	for _, record := range records {
		if len(record) > 1 && record[1] == fixletIDToDelete {
			deleted = true
			continue // Skip this record to delete it
		}
		newRecords = append(newRecords, record)
	}

	if !deleted {
		fmt.Println("FixletID not found.")
		return
	}

	newFile, err := os.Create(filepath) // Creating a new file to overwrite old content
	if err != nil {
		log.Fatal("Error creating new file: ", err)
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	writer.WriteAll(newRecords) // For Writing updated records back to file
	writer.Flush()

	fmt.Println("Record deleted successfully!")
}

func main() {
	filepath := "fixlets.csv"
	var choice int

	for {
		// The display menu options
		fmt.Println("1. List CSV file")
		fmt.Println("2. Add record to CSV")
		fmt.Println("3. Delete record by FixletID")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// Clear the newline character left in the input buffer
		bufio.NewReader(os.Stdin).ReadString('\n')

		switch choice {
		case 1:
			listCSV(filepath)
		case 2:
			addRecord(filepath)
		case 3:
			deleteRecord(filepath)
		case 4:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
