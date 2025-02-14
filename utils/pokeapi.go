package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// const for the baseline http server from where we make requests to
const baseURL string = "https://pokeapi.co/api/v2/"

// function to get the next set of 20 locations when we call map in cli
func (c *Client) getMap(pageURL *string) (locations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	// functionality to check if the getMap requests data we already have cached
	// to return instead of using the API again
	if val, ok := c.cache.Get(url); ok {
		location_response := locations{}
		err := json.Unmarshal(val, &location_response)
		if err != nil {
			return locations{}, err
		}

		return location_response, nil
	}

	// newrequest is used to "build" our request so that we can use it with c.httpClient.Do()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations{}, err
	}

	// the Do method now performs the api data request to receive the data we're looking for, and save it in resp
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locations{}, err
	}
	defer resp.Body.Close()

	// ReadAll translates the data for us that we got from the above Do request.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locations{}, err
	}

	// we create a new instance of our locations struct and we use unmarshal to unpack
	// the bytearray from data to our struct format to make it readable
	location_response := locations{}
	err = json.Unmarshal(data, &location_response)

	// debug for the json data
	// fmt.Printf("unpacked data: %v\n\n", location_response)

	if err != nil {
		return locations{}, err
	}

	c.cache.Add(url, data)
	return location_response, nil
}
