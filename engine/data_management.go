package engine

import (
	"encoding/json"
	"os"
)

// Data management behaviour
type DataManagement interface {
	LoadDataset() (DataSet, error)
	UpdateDataset() error
}

// DataManagement implementation
type DataManager struct {
	datasetPath string
}

// instantiate a new DataManager object
func NewDataManager(datasetPath string) *DataManager {
	return &DataManager{datasetPath: datasetPath}
}

// load the dataset from the storage using
// the provided path
func (uow *DataManager) LoadDataset() (*RefinedDataSet, error) {
	var dataSet DataSet

	binary, err := os.ReadFile(uow.datasetPath)
	if err != nil {
		return &RefinedDataSet{}, err
	}
	err = json.Unmarshal(binary, &dataSet)

	return dataSet.Refine(), nil
}

// update an existing dataset
func (uow *DataManager) UpdateDatatset() error {
	return nil
}
