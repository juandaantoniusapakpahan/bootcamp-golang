package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFileOs(namafile string) *os.File {
	write, err := os.Open(namafile)
	if err != nil {
		panic(err)
	}

	return write
}

func ReadFileIo(namafile string) []byte {
	bytes, err := ioutil.ReadFile(namafile)
	if err != nil {
		panic(err)
	}

	return bytes
}

type M map[string]interface{}

func Nomor1() M {
	write := ReadFileIo("soal1.json")
	result := M{}
	json.Unmarshal(write, &result)
	return result
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Location  string `json:"locations"`
	Online    bool   `json:"online"`
	Followers int    `json:"followers"`
}

type ListProfiles struct {
	Profile []Profile `json:"profile"`
}

func Nomor2() int {
	profiles := ListProfiles{}
	write := ReadFileOs("soal2.json")
	decoder := json.NewDecoder(write)
	if err := decoder.Decode(&profiles); err != nil {
		panic(err)
	}

	count := 0
	for _, v := range profiles.Profile {
		count += v.Followers
	}

	return count

}

func main() {
	nomor1 := Nomor1()
	fmt.Println(nomor1)
	nomor2 := Nomor2()
	fmt.Println(nomor2)
}
