package usecase

import (
	"osint_agent/internal/domain"
	"osint_agent/pkg/classifier"
)

type ClassificationUseCase struct {
	dataset []domain.Category
}

type TextClassification interface {
	ProcessSentences(preparedSentences []domain.Sentence) []domain.Sentence
	GroupSentences(processedSentences []domain.Sentence) error
	GetMainContext(recognizedSentences []domain.Sentence) []string
}

func NewClassificationUseCase(dataset []domain.Category) ClassificationUseCase {
	return ClassificationUseCase{
		dataset: dataset,
	}
}

func ProcessSentences(sentenceCollection []domain.Sentence) ([]domain.Sentence, error) {
	processedSentences, err := classifier.ProcessSentences(sentenceCollection)
	if err != nil {
		return nil, err
	}
	return processedSentences, nil
}

func GroupSentences(processedSentences []domain.Sentence) error {
	return nil
}

func GetMainContext(recognizedSentences []domain.Sentence) []string {
	return nil
}
