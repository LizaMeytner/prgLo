package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New(serviceName string) *Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	log := zerolog.New(output).
		With().
		Timestamp().
		Str("service", serviceName).
		Logger()

	return &Logger{log}
}

// Методы-обертки для удобства
func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.Logger.Debug().Fields(fields).Msg(msg)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.Logger.Info().Fields(fields).Msg(msg)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.Logger.Warn().Fields(fields).Msg(msg)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.Logger.Error().Fields(fields).Msg(msg)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.Logger.Fatal().Fields(fields).Msg(msg)
}
