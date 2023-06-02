package logging

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup() {
	zerolog.TimestampFieldName = "ts"
	zerolog.MessageFieldName = "msg"

	writer := zerolog.ConsoleWriter{Out: os.Stderr}
	log.Logger = zerolog.New(writer).Level(getLevel()).With().Timestamp().Caller().Stack().Logger()

	log.Debug().Msg("Logger setup complete!")
}

func getLevel() zerolog.Level {
	level := os.Getenv("LOG_LEVEL")

	switch strings.ToUpper(level) {
	case "FATAL":
		return zerolog.FatalLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "WARN":
		return zerolog.WarnLevel
	case "DEBUG", "TRACE":
		return zerolog.DebugLevel
	default:
		return zerolog.InfoLevel
	}
}
