package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vinicch/shortener-go/core/domain"
	"github.com/vinicch/shortener-go/core/port"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Uses an alias to create a shortened version of the provided URL.
// If the alias is not provided, it will generate a random one
func Shorten(createURL port.CreateURL, doesAliasExist port.DoesAliasExist,
	url, alias string) (domain.Url, error) {

	if strings.TrimSpace(alias) != "" {
		if doesAliasExist(alias) {
			return domain.Url{}, errors.New(domain.AliasAlreadyExists)
		}
	} else {
		alias = generateAlias(doesAliasExist)
	}

	host := fmt.Sprintf("http://%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	entity := domain.Url{
		Id:        uuid.NewString(),
		Alias:     alias,
		Original:  url,
		Shortened: fmt.Sprintf("%s/url/%s", host, alias),
	}

	err := createURL(&entity)
	if err != nil {
		return domain.Url{}, err
	}

	return entity, nil
}

// Generates an alias using a UTC time seed and the UTF-8 alphanumeric characters
func generateAlias(doesAliasExist port.DoesAliasExist) string {
	rand.Seed(time.Now().UTC().UnixNano())

	alias := make([]byte, 6)
	for i := range alias {
		index := rand.Intn(len(chars))
		alias[i] = chars[index]
	}

	if doesAliasExist(string(alias)) {
		return generateAlias(doesAliasExist)
	}

	return string(alias)
}
