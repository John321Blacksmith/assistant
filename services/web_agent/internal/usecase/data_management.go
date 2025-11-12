package usecase

import (
	classifierApi "osint_agent/libs/classifier/api"
)

type DataManagement interface {
	LoadDataSet(source string) ([]map[string][]string, error)
	UpdateDataSet(object map[string]map[string][]string) error
}

type DataUseCase struct {
	DatasetLocation string
}

func NewDataUseCase(datasetLocation string) *DataUseCase {
	return &DataUseCase{
		DatasetLocation: datasetLocation,
	}
}

func (uc *DataUseCase) LoadDataset() ([]map[string][]string, error) {
	categories, err := classifierApi.LoadDataset(uc.DatasetLocation)
	if err != nil {
		return nil, err
	}
	return categories.Categories, nil
}

func UpdateDataset(object map[string][]string) error {
	return nil
}
