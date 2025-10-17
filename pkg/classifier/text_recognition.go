package classifier

import (
	"tesla_go/internal/domain"
)

type TextRecognizer struct {
	knownSentences   []domain.Sentence
	unknownSentences []domain.Sentence
}

func NewTextRecognizer(
	knownSentences []domain.Sentence,
	unknownSentences []domain.Sentence,
) *TextRecognizer {
	return &TextRecognizer{
		knownSentences:   knownSentences,
		unknownSentences: unknownSentences,
	}
}

func ProcessSentences(preparedSentences []domain.Sentence) []domain.Sentence {
	return nil
}

func GroupSentences(processedSentences []domain.Sentence) error {
	return nil
}

func GetMainContext(recognizedSentences []domain.Sentence) []string {
	return nil
}
