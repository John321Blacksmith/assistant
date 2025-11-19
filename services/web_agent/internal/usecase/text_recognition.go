package usecase

import (
	classifierApi "osint_agent/libs/classifier/api"
)

type ClassificationUseCase struct {
	dataset *classifierApi.DataSet
}

type TextClassification interface {
	ProcessSentences(preparedSentences []map[string][]string) []map[string][]string
	GroupSentences(processedSentences []map[string][]string) error
	GetMainContext(recognizedSentences []map[string][]string) []string
}

func NewClassificationUseCase(dataset *classifierApi.DataSet) *ClassificationUseCase {
	return &ClassificationUseCase{
		dataset: dataset,
	}
}

func ProcessSentences(sentenceCollection []map[string][]string) ([]map[string][]string, error) {
	processedSentences, err := classifierApi.ProcessSentences(sentenceCollection)
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
