package usecase

import (
	"osint_agent/internal/domain"
	"osint_agent/pkg/classifier"
)

type DataManagement interface {
	LoadDataSet(source string) ([]domain.Category, error)
	UpdateDataSet(object map[string]domain.Category) error
}

type DataUseCase struct{}

func NewDataUseCase() DataUseCase {
	return DataUseCase{}
}

func (uc *DataUseCase) LoadDataset(source string) ([]domain.Category, error) {
	categories, err := classifier.LoadDataset(source)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func UpdateDataset(object map[string]domain.Category) error {
	return nil
}
