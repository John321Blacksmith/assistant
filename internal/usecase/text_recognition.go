package usecase

import (
	"tesla_go/internal/domain"
)

type ClassificationUseCase struct{}

type TextClassification interface {
	ProcessSentences(preparedSentences []domain.Sentence) []domain.Sentence
	GroupSentences(processedSentences []domain.Sentence) error
	GetMainContext(recognizedSentences []domain.Sentence) []string
}

func NewClassificationUseCase() ClassificationUseCase {
	return ClassificationUseCase{}
}

func ProcessSentences(sentenceCollection []domain.Sentence) []domain.Sentence {
	return nil
}

func GroupSentences(processedSentences []domain.Sentence) error {
	return nil
}

func GetMainContext(recognizedSentences []domain.Sentence) []string {
	return nil
}
