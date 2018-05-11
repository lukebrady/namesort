package main

import (
	"log"
	"os"
	"sort"
)

func main() {
	// Create the log file that will be created and appended to during runtime.
	logFile, err := os.OpenFile("./test/namesort.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	// Set the output of the logger.
	log.SetOutput(logFile)
	log.Println("Starting namesort.")

	yo, err := ScanNameFile("./test/testdns")
	if err != nil {
		panic(err)
	}
	sort.Strings(yo)
	for _, line := range yo {
		println(line)
	}
}
