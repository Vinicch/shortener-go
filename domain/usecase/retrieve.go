package usecase

import (
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/vinicch/shortener-go/domain"
	"github.com/vinicch/shortener-go/domain/port"
)

const aliasNotInformed = "'alias' parameter not informed"

// Gets an URL that corresponds to a given alias
func Retrieve(getURL port.GetURL, updateURL port.UpdateURL, alias string) (string, error) {
	if strings.TrimSpace(alias) == "" {
		return "", errors.New(aliasNotInformed)
	}

	url, err := getURL(alias)
	if err != nil {
		log.Error().Err(err).Msg("Error getting URL info")
		return "", err
	} else if url == nil {
		return "", errors.New(domain.ShortenedURLNotFound)
	}

	url.Visits++

	go updateURL(url)

	return url.Original, nil
}
