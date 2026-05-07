package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const LOCATION_AREAS_URL = "https://pokeapi.co/api/v2/location-area"

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url string) (LocationAreasResponse, error) {
	if cached, exists := c.cache.Get(url); exists {
		var cachedRes LocationAreasResponse
		if err := json.Unmarshal(cached, &cachedRes); err == nil {
			return cachedRes, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error fetching map locations: %w", err)
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

	c.cache.Add(url, dat)

	return parsedRes, nil
}
