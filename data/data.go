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

func LoadWebLogData() []WebLog {
	data := readCsvFile("data/weblog.csv")
	webLogs := make([]WebLog, len(data)-1)
	j := 0
	for i := 1; i < len(data); i++ {
		row := data[i]
		webLog := WebLog{IP: row[0], Time: row[1], URL: row[2], Status: row[3]}
		webLogs[j] = webLog
		j++
	}
	return webLogs
}
