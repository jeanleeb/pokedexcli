package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetAreaDetails(area string) (AreaResponse, error) {
	if area == "" {
		return AreaResponse{}, errors.New("invalid area")
	}

	url := LOCATION_AREAS_URL + "/" + area

	if cached, exists := c.cache.Get(url); exists {
		var cachedRes AreaResponse
		if err := json.Unmarshal(cached, &cachedRes); err == nil {
			return cachedRes, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return AreaResponse{}, fmt.Errorf("error fetching area details: %w", err)
	}
	if res.StatusCode == http.StatusNotFound {
		return AreaResponse{}, fmt.Errorf("%s not found", area)
	}
	if res.StatusCode != http.StatusOK {
		return AreaResponse{}, fmt.Errorf("unexpected response status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaResponse{}, fmt.Errorf("error reading respone body: %w", err)
	}

	var parsedRes AreaResponse
	if err = json.Unmarshal(dat, &parsedRes); err != nil {
		return AreaResponse{}, fmt.Errorf("error parsing response: %w", err)
	}

	c.cache.Add(url, dat)

	return parsedRes, nil
}
