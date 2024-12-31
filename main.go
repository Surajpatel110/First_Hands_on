package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func listCSV(file *os.File) {
	reader := csv.NewReader(file)

	//Now after creating the reader we'll read all the content in csv file

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error in Reading the csv file %v/n", err)
		return
	}
	//Printing the records
	for i, record := range records {
		fmt.Printf("Row %d: %v\n", i+1, record)
	}

}

func AddRecord(file *os.File) {

	f, err := os.OpenFile("fixlets.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	fmt.Println("writing into a file")
	writer := csv.NewWriter(f)

	arr := []string{"1fdsfsfdsf", "50121701000", "MS22-AUG: Security Update for Windows Server 2022 - Windows Server 2022 - KB50121701000 (x64)", "Critical", "96"}
	fmt.Print("writting to file")
	err1 := writer.Write(arr)

	if err1 != nil {
		log.Fatal("there is an error while writing to file ", err1.Error())
	}
	writer.Flush()
	if err2 := writer.Error(); err2 != nil {
		log.Fatal("there is some error", err2.Error())
	}
	fmt.Println("file written succesfully")
}

func deleteRecord(file *os.File) {
	reader := csv.NewReader(file)
	records, err1 := reader.ReadAll()

	if err1 != nil {
		log.Fatal(err1.Error())
	}
	//this will create a file or override the existing file
	newFile, err := os.Create("fixlets.csv")
	if err != nil {
		log.Fatal("there is some error ", err.Error())
	}

	writer := csv.NewWriter(newFile)

	for indx, record := range records {
		if indx > len(records)-2 {
			break
		}
		err := writer.Write(record)

		if err != nil {
			log.Fatal(err.Error())
		}

	}
	writer.Flush()
	fmt.Println("<<<<<<<<<<the last record has been deleted>>>>>>>>>")

}
func main() {
	filepath := "fixlets.csv"

	//Open the csv file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error in Opening csv File %v/n", err)
		return
	}
	defer file.Close()
	var choice int

	for {
		fmt.Print("1. List Csv file \n 2.Add record to csv \n 3. Delete Record from csv \n")
		fmt.Print("enter your choice ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			listCSV(file)

		case 2:
			AddRecord(file)

		case 3:
			deleteRecord(file)

		default:
			fmt.Print("wrong choice ")

		}
		if choice > 3 {
			break
		}
	}

}
