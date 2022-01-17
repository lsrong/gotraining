package main

// 演示如何自定义json序列化和反序列化方法

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

// UnmarshalJSON 自定义json反序列化接口方法, implement json.Unmarshaler
func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	default:
		*a = Unknown
	}

	return nil
}

// MarshalJSON 自定义
func (a *Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch *a {
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	default:
		s = "Unknown"
	}

	return json.Marshal(s)
}

func main() {
	// 反序列化animal的json字符串
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(zoo)

	// 简单计数
	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal]++
	}

	fmt.Printf(`Zoo Cencus:
*Gopher: %d
*Zebras:%d
*Unknown:%d
`, census[Gopher], census[Zebra], census[Unknown])

}
