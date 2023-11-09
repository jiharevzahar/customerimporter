// package main reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package importer

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CustomerImporter struct {
	path string
}

func NewCustomerImporter(path string) CustomerImporter {
	return CustomerImporter{path: path}
}

func (c *CustomerImporter) Import() ([][]string, error) {
	file, err := os.Open(c.path)
	if err != nil {
		return nil, fmt.Errorf("importer: cannot open file: %w", err)
	}

	parser := csv.NewReader(file)
	records, err := parser.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("importer: could not read file: %w", err)
	}
	return records, nil
}
