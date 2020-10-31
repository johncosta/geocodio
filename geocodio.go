package geocodio

import (
	"errors"
	"net/http"
	"time"
)

const (
	// GeocodioAPIBaseURLv1 is the Geocod.io Base URL
	GeocodioAPIBaseURLv1 = "https://api.geocod.io/v1.4"
)

// NewGeocodio is a helper to create new Geocodio pointer
func NewGeocodio(apiKey string) (*Geocodio, error) {

	if apiKey == "" {
		return nil, errors.New("apiKey is missing")
	}

	newGeocodio := new(Geocodio)
	newGeocodio.APIKey = apiKey
	newGeocodio.client = &http.Client{
		Timeout: 10 * time.Second,
	}
	return newGeocodio, nil
}

// NewGeocodioWithClient is a helper to create new Geocodio pointer
func NewGeocodioWithClient(apiKey string, client *http.Client) (*Geocodio, error) {
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	newGeocodio, err := NewGeocodio(apiKey)
	newGeocodio.client = client

	return newGeocodio, err
}
