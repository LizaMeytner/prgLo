package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Log = zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Для production можно настроить уровень логирования
	// Log = Log.Level(zerolog.InfoLevel)
}
