package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

func main() {
	//Base url
	url :=  "https://api.weather.gov"
		
	//User prompt stuff here to get values
	var x, y string
	fmt.Println("Enter x and y coordinates of the location you would like to see the weather of, as integers:")
	fmt.Scan(&x, &y)

	weather := url + "/gridpoints/TOP/" + x + "," + y + "/forecast"
	fmt.Println(weather)

	//Simple get request
	resp, err := http.Get(weather)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return
	}

	//If it got here then the request was successful so read the values
	jsonData, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

	var data map[string]interface{}
	
	err2 := json.Unmarshal([]byte(jsonData), &data)
    if err2 != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

	//Filter through json data to be more readable
    properties := data["properties"].(map[string]interface{})
    periods := properties["periods"].([]interface{})
    firstPeriod := periods[0].(map[string]interface{})
    description := firstPeriod["detailedForecast"].(string)
	fmt.Println("At the location you selected, here is the forcast:", description)
}