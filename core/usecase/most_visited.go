package usecase

import (
	"github.com/vinicch/shortener-go/core/domain"
	"github.com/vinicch/shortener-go/core/port"
)

// Gets the 10 most visited URLs
func MostVisited(getMostVisited port.GetMostVisited) ([]domain.Url, error) {
	return getMostVisited()
}
