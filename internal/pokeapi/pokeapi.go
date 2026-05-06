package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreasResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("Error fetching map locations: %w", err)
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error reading respone body: %w", err)
	}

	var parsedRes LocationAreasResponse
	if err = json.Unmarshal(dat, &parsedRes); err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error parsing response: %w", err)
	}

	return parsedRes, nil
}
