package data

import (
	"encoding/csv"
	"log"
	"os"
)

type WebLog struct {
	IP     string
	Time   string
	URL    string
	Status string
}

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

func ReadTweets(filPath string) []string {
	rows := readCsvFile(filPath)
	tweets := make([]string, len(rows))
	for i := range rows {
		tweets[i] = rows[i][5]
	}
	return tweets
}
