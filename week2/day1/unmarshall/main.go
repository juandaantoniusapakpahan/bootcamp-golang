package main

import (
	"encoding/json"
	"fmt"
)

type Alamat struct {
	Alamat string `json:"alamat"`
}

type EmbedStruct struct {
	Name   string `json:"name"`
	Alamat `json:"address"`
}

type NestedStruct struct {
	Name    string `json:"name"`
	Address Alamat `json:"address"`
}

type M map[string]interface{}

func main() {
	jsonData := `{"name":"Juanda","address":{"alamat":"Medan"}}`

	nestedData := NestedStruct{}
	err := json.Unmarshal([]byte(jsonData), &nestedData)
	if err != nil {
		panic(err)
	}

	embedData := EmbedStruct{}
	err = json.Unmarshal([]byte(jsonData), &embedData)
	if err != nil {
		panic(err)
	}

	mapData := M{}
	err = json.Unmarshal([]byte(jsonData), &mapData)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Nested Test:", nestedData)
	fmt.Println()
	fmt.Println()
	fmt.Println("Embed Test:", nestedData)
	fmt.Println()
	fmt.Println()
	fmt.Println("Map Test:", mapData)
	fmt.Println()

}
