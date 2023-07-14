package main

import (
	"encoding/json"
	"fmt"
)

type Alamat struct {
	Alamat string `json:"alamat"`
}

type NestedStruct struct {
	Name    string `json:"name"`
	Address Alamat `json:"address"`
}

type EmbedStruct struct {
	Name string `json:"name"`
	Alamat
}

type M map[string]interface{}

func main() {
	testNested := NestedStruct{
		Name:    "Nested",
		Address: Alamat{"Rantauprapat"},
	}

	testEmbed := EmbedStruct{
		Name:   "Embed",
		Alamat: Alamat{"GGWP"},
	}

	testMap := M{"Name": "Map", "Alamat": "Medan dan sekitar"}

	testString := `
	{
		"name":"juanda",
		"umur":21
	}
	`

	nestedByte, err := json.Marshal(testNested)
	if err != nil {
		panic(err)
	}

	embedByte, err := json.Marshal(testEmbed)
	if err != nil {
		panic(err)
	}

	mapByte, err := json.Marshal(testMap)
	if err != nil {
		panic(err)
	}

	strintByte, err := json.Marshal(json.RawMessage(testString))

	fmt.Println()
	fmt.Println("Nested Struct Byte:", nestedByte)
	fmt.Println("Nested Struct Json:", string(nestedByte))
	fmt.Println()
	fmt.Println("Embed Struct Byte:", embedByte)
	fmt.Println("Embed Struct Json:", string(embedByte))
	fmt.Println()
	fmt.Println("Map Byte:", mapByte)
	fmt.Println("Map Json:", string(mapByte))
	fmt.Println()
	fmt.Println("String Byte:", strintByte)
	fmt.Println("String Json:", string(strintByte))
	fmt.Println()
}
