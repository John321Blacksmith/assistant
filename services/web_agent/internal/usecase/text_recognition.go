package usecase

import (
	"osint_agent/libs/classifier/api"
)

type ClassificationUseCase struct {
	dataset []map[string][]string
}

type TextClassification interface {
	ProcessSentences(preparedSentences []map[string][]string) []map[string][]string
	GroupSentences(processedSentences []map[string][]string) error
	GetMainContext(recognizedSentences []map[string][]string) []string
}

func NewClassificationUseCase(dataset []map[string][]string) *ClassificationUseCase {
	return &ClassificationUseCase{
		dataset: dataset,
	}
}

func ProcessSentences(sentenceCollection []map[string][]string) ([]map[string][]string, error) {
	processedSentences, err := api.ProcessSentences(sentenceCollection)
	if err != nil {
		return nil, err
	}
	return processedSentences, nil
}

func GroupSentences(processedSentences []map[string][]string) error {
	return nil
}

func GetMainContext(recognizedSentences []map[string][]string) []string {
	return nil
}
