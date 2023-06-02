package port

import "github.com/vinicch/shortener-go/core/domain"

// Gets information about an URL from the repository
type GetURL func(string) (*domain.Url, error)

// Gets the 10 most visited URLs
type GetMostVisited func() ([]domain.Url, error)

// Creates an URL record containing information about it and its alias
type CreateURL func(*domain.Url) error

// Updates URL information
type UpdateURL func(*domain.Url)

// Checks if a record for the given alias already exists
type DoesAliasExist func(string) bool

// These abstractions can be an interface as well (might be better).
type Repository interface {
	// Gets information about an URL from the repository
	GetURL(string) (*domain.Url, error)

	// Gets the 10 most visited URLs
	GetMostVisited() ([]domain.Url, error)

	// Creates an URL record containing information about it and its alias
	CreateURL(*domain.Url) error

	// Updates URL information
	UpdateURL(*domain.Url)

	// Checks if a record for the given alias already exists
	DoesAliasExist(string) bool
}
