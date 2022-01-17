// Sample program to show how to marshal a user defined
// struct type into a string.
// 演示如何将用户定义的结构类型编组为字符串的示例程序。
package main

import (
	"encoding/json"
	"fmt"
)

type (
	Condition struct {
		WindSpeed     float64 `json:"wind_speed_milehour"`
		WindDirection int     `json:"wind_direction_degnorth"`
		GustWindSpeed float64 `json:"gust_wind_speed_milehour"`
	}
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}
	// BuoyStation 站点信息
	BuoyStation struct {
		StationID string    `json:"station_id"`
		Name      string    `json:"name"`
		LocDesc   string    `json:"location_desc"`
		Page      string    `json:"station_page"`
		Condition Condition `json:"condition"`
		Location  Location  `json:"location"`
	}
)

// {StationID:sblf1 Name:Station SBLF1 - 8726673 LocDesc:Seabulk, Tampa, FL Page:http://www.ndbc.noaa.gov/station_page.php?station=sblf1 Condition:{WindSpeed:4.697573786668777 WindDirection:350 GustWindSpeed:5.816043786668778} Location:{Type:Point Coordinates:[-82.445 27.923]}}

func main() {
	station := BuoyStation{
		StationID: "sblf1",
		Name:      "Station SBLF1 - 8726673",
		LocDesc:   "Seabulk, Tampa, FL",
		Page:      "http://www.ndbc.noaa.gov/station_page.php?station=sblf1",
		Condition: Condition{
			WindSpeed:     4.697573786668777,
			WindDirection: 350,
			GustWindSpeed: 5.816043786668778,
		},
		Location: Location{
			Type:        "Point",
			Coordinates: []float64{-82.445, 27.923},
		},
	}

	// json.Marshal, encode to json string
	b, err := json.Marshal(&station)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("%s \n", b)
	fmt.Println("-------------------------")

	// json.MarshalIndent:
	// Each JSON element in the output will begin on a new line beginning with prefix followed by one
	// or more copies of indent according to the indentation nesting
	b2, err := json.MarshalIndent(&station, "|", "--")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf("%s \n", b2)
}
