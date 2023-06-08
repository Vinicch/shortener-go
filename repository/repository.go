package repository

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/vinicch/shortener-go/domain/port"
)

type urlFunctions struct {
	port.GetURL
	port.GetMostVisited
	port.CreateURL
	port.UpdateURL
	port.DoesAliasExist
}

func MakeURLFunctions() urlFunctions {
	db := os.Getenv("DATABASE")

	switch strings.ToLower(db) {
	case "postgres":
		conn := createPgConnection()

		return urlFunctions{
			getURL(conn),
			getMostVisited(conn),
			createURL(conn),
			updateURL(conn),
			doesAliasExist(conn),
		}
	}

	log.Fatal().Str("database", db).Msg("There's no implementation for this database")
	return urlFunctions{}
}
