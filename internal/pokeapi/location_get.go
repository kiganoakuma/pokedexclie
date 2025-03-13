package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationRes := Location{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Location{}, err
		}

		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}
	locationRes := Location{}
	err = json.Unmarshal(dat, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return locationRes, nil
}
