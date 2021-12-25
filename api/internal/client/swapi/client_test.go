package swapi_test

import (
	"github.com/streetbyters/aduket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func generateStarshipDetailResponse() *StarshipDetailResponse {
	return &StarshipDetailResponse{
		Name:                 "Death Star",
		Model:                "DS-1 Orbital Battle Station",
		Manufacturer:         "Imperial Department of Military Research, Sienar Fleet Systems",
		CostInCredits:        "1000000000000",
		Length:               "120000",
		MaxAtmospheringSpeed: "n/a",
		Crew:                 "342,953",
		Passengers:           "843,342",
		CargoCapacity:        "1000000000000",
		Consumables:          "3 years",
		HyperdriveRating:     "4.0",
		Mglt:                 "10",
		StarshipClass:        "Deep Space Mobile Battlestation",
		Pilots:               nil,
		Films:                nil,
		Created:              time.Time{},
		Edited:               time.Time{},
		URL:                  "https://swapi.dev/api/starships/9/",
	}
}
func TestGivenCorrectAndExistStarshipNameWhenGetStarshipDetailsCalledThenItShouldReturnExpectedDetails(t *testing.T) {
	expectedStarshipDetails := generateStarshipDetailResponse()
	givenStarshipName := StarshipDetailRequest{Name: expectedStarshipDetails.Name}

	testServer, _ := aduket.NewServer(
		http.MethodGet,
		"/api/starships/9/",
		aduket.JSONBody(expectedStarshipDetails),
		aduket.StatusCode(200))
	defer testServer.Close()

	client := NewSwapiClient("/api/starships/9/")
	starshipDetails, err := client.GetStarshipDetails(givenStarshipName.Name)

	assert.Nil(t, err)
	assert.Equal(t, expectedStarshipDetails, starshipDetails)
}
