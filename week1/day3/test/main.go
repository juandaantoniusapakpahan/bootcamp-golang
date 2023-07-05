package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Voco(names []string) {

	vocal := "aiuoe"

	for i := 0; i < len(names); i++ {
		count := 0
		for j := 0; j < len(names[i]); j++ {
			res1 := strings.ToLower(string(names[i][j]))

			if strings.Contains(vocal, res1) {
				count++
			}
		}
		fmt.Println(names[i], ":", count, "huruf vocal")

	}
}

func GetNameByPartName(names []string, name string) []string {

	nameLower := strings.ToLower(name)
	var result []string
	for i := 0; i < len(names); i++ {
		res1 := strings.ToLower(names[i])
		if strings.Contains(res1, nameLower) {
			result = append(result, names[i])
		}
	}
	return result
}

type M map[string]int
type DataBosQ struct {
	Vocal map[string]int
}

// NO # 3

func (d *DataBosQ) CountVocName(names []string) {

	for i := 0; i < len(names); i++ {
		char := strings.ToLower(string(names[i][0]))
		vocal := "aiueo"
		if strings.Contains(vocal, char) {

			if len(d.Vocal) == 0 {
				d.Vocal = M{char: 1}
				continue
			}
			val, ok := d.Vocal[char]
			if ok {
				d.Vocal[char] = val + 1
			} else {
				d.Vocal[char] = 1
			}

		}

	}

}

func GetInput() []string {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	str := []string{}

	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	return str
}

func main() {

	slice := GetInput()

	// NO 1
	Voco(slice)
	fmt.Println()

	// NO 2
	result := GetNameByPartName(slice, "Evans")
	fmt.Println(result)
	fmt.Println()

	// NO 3
	data := DataBosQ{}
	data.CountVocName(slice)
	fmt.Println(data.Vocal)
}
