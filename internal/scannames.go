package main

import (
	"bufio"
	"os"
	"regexp"
)

// ScanNameFile scans in a given name file and returns
// all name entries within the file.
func ScanNameFile(file string) ([]string, error) {
	var names []string
	isNameSort := regexp.MustCompile("A|CNAME")
	// Open name file for reading.
	nameFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	// Defer closing of the name file.
	defer nameFile.Close()
	// Create a new reader for the nameFile.
	scanner := bufio.NewScanner(nameFile)
	// Iterate over the file until io.EOF.
	for scanner.Scan() {
		// Check to see if the entry is an A record or CNAME.
		if isNameSort.MatchString(scanner.Text()) {
			names = append(names, scanner.Text())
		}
	}
	// Check to see if any errors occur.
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return names, nil
}
