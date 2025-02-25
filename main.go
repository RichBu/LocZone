package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const INPUT_FILE = "ZIP_Locale_Detail_CSV.csv"

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	records := readCsvFile(INPUT_FILE)
	for _, record := range records {
		distNo := record[3]
		zipCode := record[4]
		fmt.Println(distNo, zipCode)
		//fmt.Println(record)
	}
	//fmt.Println(records)
}
