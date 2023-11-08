package logger

import (
	"errors"
	"fmt"
	"ipr-savelichev/internal/config"

	"go.uber.org/zap"
)

const (
	debugLevel   = "debug"
	releaseLevel = "release"
)

var (
	ErrUnknowParam = errors.New("Unknow parametr")
)

func NewLogger(config *config.Config) (*zap.Logger, error) {
	var (
		err    error
		logger *zap.Logger
	)

	switch config.LogLevel {
	case debugLevel:
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		logger.Info("Logger debug level is on.")
	case releaseLevel:
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
		logger.Info("Logger info level is on.")
	default:
		return nil, fmt.Errorf("%w %v", ErrUnknowParam, config.LogLevel)
	}

	return logger, err
}
