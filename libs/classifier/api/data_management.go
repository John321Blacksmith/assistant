package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataSet struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

func LoadDataset(filename string) (*DataSet, error) {
	var dataset *DataSet
	file_bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error occurred processing the file %v\n", filename)
	} else {
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
