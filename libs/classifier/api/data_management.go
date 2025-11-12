package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataSet struct {
	Categories []map[string][]string `json:"categories"`
}

func LoadDataset(filename string) (*DataSet, error) {
	var categories []map[string][]string
	// form contents of the file to bytes
	file_bytes, err := os.ReadFile(filename)
	fmt.Println(categories)
	if err != nil {
		fmt.Printf("Error occurred processing the file %v\n", filename)
	} else {
		// transform the data from bytes to the datatype
		err = json.Unmarshal(file_bytes, &categories)
		if err != nil {
			fmt.Printf("Error while reading the file %v, Error: %v\n", filename, err)
		} else {
			return &DataSet{Categories: categories}, nil
		}
	}
	return nil, err
}

func UpdateDataSet(object map[string][]string) error {
	return nil
}
