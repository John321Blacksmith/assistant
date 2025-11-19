// Package domain provides business
// oriented DTOs.

package domain

import (
	"time"
)

// Struct Sentence represents
// a sentence entity with its
// defined category and contents.
type Sentence struct {
	Category string
	Data     []string
}

// Struct Config represents an
// entity of the main cofig for
// application life.
type Config struct {
	Environment string        `json:"environment"`
	Parsing     ParsingConfig `json:"parsing"`
	Http        HTTP          `json:"http"`
	Prometheus  PromConfig    `json:"prometheus"`
}

// Struct HTTP represents settings
// for an HTTP server startup.
type HTTP struct {
	Host    string        `json:"host"`
	Port    string        `json:"port"`
	Timeout time.Duration `json:"duration"`
}

// Struct HTTP represents settings
// for facilitation of sending logs
// to the Prometheus sevrer.
type PromConfig struct {
	ConsumerAddr string `json:"promAddress"`
}

// Struct HTTP represents settings
// for TextClassifier.
type ParsingConfig struct {
	DatasetLocation string `json:"datasetLocation"`
}
