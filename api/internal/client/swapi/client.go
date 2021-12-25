package swapi

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Swapi struct {
	URL  string
	http *http.Client
}

func NewSwapiClient(url string) *Swapi {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		},
	}
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	return &Swapi{
		URL:  url,
		http: &client,
	}
}

func (s *Swapi) GetStarshipDetails(starshipName string) (*StarshipDetailResponse, error) {
	panic("implement me")
}
