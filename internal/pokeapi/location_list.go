package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsRes := ResShallowLocations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return ResShallowLocations{}, err
		}

		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResShallowLocations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return ResShallowLocations{}, err
	}
	locationsRes := ResShallowLocations{}
	err = json.Unmarshal(dat, &locationsRes)
	if err != nil {
		return ResShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsRes, nil
}
