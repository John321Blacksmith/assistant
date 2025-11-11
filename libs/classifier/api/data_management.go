package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadDataset(filename string) ([]map[string][]string, error) {
	var dataset []map[string][]string
	// form contents of the file to bytes
	file_bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error occurred processiong the file %v\n", filename)
	} else {
		// transform the data from bytes to the datatype
		err = json.Unmarshal(file_bytes, &dataset)
		if err != nil {
			fmt.Printf("Error while reading the file %v, Error: %v\n", filename, err)
		} else {
			return dataset, nil
		}
	}
	return nil, err
}

func UpdateDataSet(object map[string][]string) error {
	return nil
}
