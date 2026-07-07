package api

type TextRecognizer struct {
	knownSentences   []map[string][]string
	unknownSentences []map[string][]string
}

func NewTextRecognizer(
	knownSentences []map[string][]string,
	unknownSentences []map[string][]string,
) *TextRecognizer {
	return &TextRecognizer{
		knownSentences:   knownSentences,
		unknownSentences: unknownSentences,
	}
}

func ProcessSentences(preparedSentences []map[string][]string) ([]map[string][]string, error) {
	return nil, nil
}

func GroupSentences(processedSentences []map[string][]string) error {
	return nil
}

func GetMainContext(recognizedSentences []map[string][]string) []string {
	return nil
}
