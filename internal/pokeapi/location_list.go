package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, found := c.cache.Get(url); found {
		// Data was found in the cache, unmarshal and return
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
