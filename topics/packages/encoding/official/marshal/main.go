package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Colors []string

type ColorGroup struct {
	ID   int
	Name string `json:"name,omitempty"`
	Colors
}

func main() {
	group := ColorGroup{
		ID:     1,
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(&group)
	if err != nil {
		fmt.Println("Error", err)
	}
	os.Stdout.Write(b)
}
