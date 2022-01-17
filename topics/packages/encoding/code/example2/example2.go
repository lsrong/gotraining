// Sample program to show how to unmarshal a JSON document into
// a user defined struct type from a file.
// 演示如何将 JSON 文档从文件解码为用户定义的结构类型的示例程序
package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	BuoyStation struct {
		StationID string    `json:"station_id"`
		Name      string    `json:"name"`
		LocDesc   string    `json:"location_desc"`
		Page      string    `json:"station_page"`
		Condition Condition `json:"condition"`
		Location  Location  `json:"location"`
	}
)

func main() {
	// open data.json file
	stream, err := os.Open("data.json")
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer stream.Close()

	// Decode the file into a slice of buoy stations
	var stations []BuoyStation
	if err = json.NewDecoder(stream).Decode(&stations); err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	// Iterate over the stations slice
	for _, station := range stations {
		fmt.Printf("%+v\n\n", station)
	}
}
