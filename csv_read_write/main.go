package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// data

func getData() *[][]string {
	return &[][]string{
		{"Name", "City", "Skills"},
		{"Smith", "Newyork", "Java"},
		{"William", "Paris", "Golang"},
		{"Rose", "London", "PHP"},
	}
}
func createCsvFile() *os.File {
	csvFile, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("Failed to creating the file. %s", err)
	}
	defer csvFile.Close()
	return csvFile
}

func readCsv(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error occured while opening the file %s", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.Read()
	if err != nil {
		log.Fatalf("Error occured during file reading %s", err)
	}
	fmt.Println(data)
}

// here f with be the csv file we want to write in
func writeCsv(filePath string, data *[][]string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644) // os.O_TRUNC
	if err != nil {
		log.Fatalf("error occured while opening the file %s", err)
	}
	// close the file to prevent memory leak
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	for _, row := range *data {
		_ = csvWriter.Write(row)
	}
	// to remove buffer from the memory
	csvWriter.Flush()
}

func main() {
	//// file := createCsvFile()
	data := getData()
	writeCsv("data.csv", data)
	readCsv("data.csv")

	//csvFile, err := os.Create("data.csv")
	//if err != nil {
	//	log.Fatalf("Failed to creating the file. %s", err)
	//}
	//defer csvFile.Close()
	////writeCsv(file, data)
	//file, err := os.Open("data.csv")
	//csvWriter := csv.NewWriter(file)
	//for _, row := range *data {
	//	_ = csvWriter.Write(row)
	//}
	//// to remove buffer from the memory
	//csvWriter.Flush()
	//
	////readCsv(csvFile)
	//csvReader := csv.NewReader(file)
	//data1, err := csvReader.ReadAll()
	//if err == nil {
	//	log.Fatalf("Error occured during file reading %s", err)
	//}
	//fmt.Println(data1)
	//file.Close()
	//create file
	//file, err := os.Open("data.csv")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//file.Close()
	////write
	//empData := [][]string{
	//	{"Name", "City", "Skills"},
	//	{"Smith", "New York", "Java"},
	//	{"William", "Paris", "Golang"},
	//	{"Rose", "London", "PHP"},
	//}
	//csvWriter := csv.NewWriter(file)
	//
	//for _, empRow := range empData {
	//	if err := csvWriter.Write(empRow); err != nil {
	//		log.Fatalf("error writing CSV: %s", err)
	//	}
	//}
	//
	//csvWriter.Flush()
	//// read
	//reader := csv.NewReader(file)
	//records, _ := reader.ReadAll()
	//fmt.Println(records)
	//fmt.Println(records[0][0])

}
