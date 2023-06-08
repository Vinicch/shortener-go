package usecase

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinicch/shortener-go/domain"
)

func TestShorten(t *testing.T) {
	host := "hostname"
	port := "80"
	url := "http://hostname.com/long/url/path"
	alias := "test"
	shortened := fmt.Sprintf("http://%s:%s/url/%s", host, port, alias)

	os.Setenv("HOST", host)
	os.Setenv("PORT", port)

	createURL := func(*domain.Url) error { return nil }
	doesAliasExist := func(string) bool { return false }

	result, err := Shorten(createURL, doesAliasExist, url, alias)

	if assert.NoError(t, err) {
		assert.Equal(t, result.Alias, alias)
		assert.Equal(t, result.Original, url)
		assert.Equal(t, result.Shortened, shortened)
	}
}

func TestShortenExistingAlias(t *testing.T) {
	alias := "test"
	createURL := func(*domain.Url) error { return nil }
	doesAliasExist := func(string) bool { return true }

	_, err := Shorten(createURL, doesAliasExist, "", alias)

	assert.EqualError(t, err, domain.AliasAlreadyExists)
}

func TestShortenDeps(t *testing.T) {
	alias := "test"
	createURLErr := "createURLError"
	createURL := func(*domain.Url) error { return errors.New(createURLErr) }
	doesAliasExist := func(string) bool { return false }

	_, err := Shorten(createURL, doesAliasExist, "", alias)

	assert.EqualError(t, err, createURLErr)
}
