// Package engine provides logic
// for data recognition implementation
// and server startup.
package engine

import (
	"fmt"
	"log/slog"
	"net/http"
	localLoggingApi "osint_agent/libs/logging/local/pkg"
	promLoggingApi "osint_agent/libs/logging/prometheus/pkg"
	"osint_agent/services/web_agent/internal/domain"
	"osint_agent/services/web_agent/internal/usecase"
)

// Engine struct represents
// an entity of the application
// body with all instances invoked.
type Engine struct {
	srv            *http.Server
	business       Business
	logger         *slog.Logger
	externalLogger *slog.Logger
}

// Struct Business represents
// a set of the usecases invoked
// in the text processing logic.
type Business struct {
	dataUseCase            *usecase.DataUseCase
	inputUseCase           *usecase.InputUseCase
	textRecognitionUseCase *usecase.ClassificationUseCase
}

// Funtion NewEngine prepares all
// instances to be invoked for the
// application lifetime.
func NewEngine(cfg *domain.Config) (*Engine, error) {
	logger := localLoggingApi.NewLogger()
	promLogger := promLoggingApi.NewLogger()
	dataUseCase := usecase.NewDataUseCase(cfg.Parsing.DatasetLocation)
	dataSet, err := dataUseCase.LoadDataset()
	if err != nil {
		logger.Error(fmt.Sprintf("Error loading the dataset: %v", err))
	}
	inputUseCase := usecase.NewInputUseCase()
	textRecognitionUseCase := usecase.NewClassificationUseCase(dataSet)

	return &Engine{
		srv: &http.Server{
			Addr: cfg.Http.Host + ":" + cfg.Http.Port,
		},
		business: Business{
			dataUseCase:            dataUseCase,
			inputUseCase:           inputUseCase,
			textRecognitionUseCase: textRecognitionUseCase,
		},
		logger:         logger,
		externalLogger: promLogger,
	}, nil
}

// Function Run starts an entire
// engine with all configs loaded.
func (app *Engine) Run() chan error {
	errChan := make(chan error)
	app.logger.Info("Starting the HTTP server...")
	if err := app.srv.ListenAndServe(); err != nil {
		app.logger.Error(fmt.Sprintf("Error occurred: %v", err))
		errChan <- err
	}
	return errChan
}
