package usecase

import (
	"osint_agent/libs/classifier/api"
)

type InputManagement interface {
	SplitToParticles(data string, sym string) []string
	RefineSentences(rawSentences []string) [][]string
	AggregateSentences(refinedSentences [][]string) []map[string][]string
}

type InputUseCase struct{}

func NewInputUseCase() *InputUseCase {
	return &InputUseCase{}
}

func (uc *InputUseCase) SplitToParticles(data string, sym string) []string {
	return api.SplitToParticles(data, sym)
}

func (uc *InputUseCase) RefineSentences(rawSentences []string) [][]string {
	return api.RefineSentences(rawSentences)
}

func (uc *InputUseCase) AggregateSentences(refinedSentences [][]string) []map[string][]string {
	return api.AggregateSentences(refinedSentences)
}
