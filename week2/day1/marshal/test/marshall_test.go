package test

import (
	"encoding/json"
	"testing"
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

func Marshall() error {
	nestedStruct := NestedStruct{
		Name:    "Juanda",
		Address: Alamat{Alamat: "Medan"},
	}

	_, err := json.Marshal(nestedStruct)
	if err != nil {
		return err
	}
	return nil
}

func BenchmarkMarshall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := Marshall()
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
}
