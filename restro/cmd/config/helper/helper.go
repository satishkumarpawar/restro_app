package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Geocode struct {
	/*Input struct {
		AddressComponents struct {
			Number          string `json:"number"`
			Street          string `json:"street"`
			Suffix          string `json:"suffix"`
			FormattedStreet string `json:"formatted_street"`
			City            string `json:"city"`
			State           string `json:"state"`
			Zip             string `json:"zip"`
			Country         string `json:"country"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
	} `json:"input"`*/
	Results []struct {
		/*AddressComponents struct {
			Number          string `json:"number"`
			Street          string `json:"street"`
			Suffix          string `json:"suffix"`
			FormattedStreet string `json:"formatted_street"`
			City            string `json:"city"`
			County          string `json:"county"`
			State           string `json:"state"`
			Zip             string `json:"zip"`
			Country         string `json:"country"`
		} `json:"address_components"`*/
		FormattedAddress string `json:"formatted_address"`
		Location         struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		/*Accuracy     int    `json:"accuracy"`
		AccuracyType string `json:"accuracy_type"`
		Source       string `json:"source"`*/
	} `json:"results"`
}

type Geolocation struct {
	Address  string `json:"address"`
	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
}

func GetGeocode(address string) Geolocation {
	key := "500a28d20d217e7e7ef51e7d5e7d7e18aadd8de"
	address = strings.Replace(address, "/", " ", -1)
	address = url.QueryEscape(address)
	//address := "2681+Overlook+Point+%2C+Escondido%2C+CA+92029"
	resp, err := http.Get("https://api.geocod.io/v1.7/geocode?q=" + address + "&api_key=" + key)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	jsonGeoData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(string(jsonGeoData))
	//return ""
	data := Geocode{}

	//NOTE : unmarshal to struct will throw this error message
	//        json: cannot unmarshal object into Go value of type

	// Decode JSON into our map

	//json.Unmarshal(byteValue, &users)
	err = json.Unmarshal([]byte(jsonGeoData), &data)
	geolocation := Geolocation{}
	if err != nil {
		println(err)
		return geolocation
	}
	//fmt.Println(data)
	for _, v := range data.Results {
		geolocation.Address = v.FormattedAddress
		geolocation.Location.Lat = v.Location.Lat
		geolocation.Location.Lng = v.Location.Lng
		//fmt.Println("Key :", k, "Lat :", v.Location.Lat, "Long", v.Location.Lng)
		break
	}

	return geolocation

}
