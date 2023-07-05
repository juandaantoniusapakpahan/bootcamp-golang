package main

import "fmt"

type List struct {
	Data []int
}

type Interface1 interface {
	Add(a int)
	Update(from, to int)
	DeleteOne(a int)
}

func (l *List) Add(a int) {
	l.Data = append(l.Data, a)

}

func (l *List) Update(from, to int) {

	for i, val := range l.Data {
		if val == from {
			l.Data[i] = to
		}
	}

}

func (l *List) DeleteOne(a int) {
	for i := 0; i < len(l.Data); i++ {
		if l.Data[i] == a {
			if i == len(l.Data)-1 {
				l.Data = l.Data[:len(l.Data)-1]
			} else {
				l.Data = append(l.Data[:i], l.Data[i+1:]...)
			}

		}
	}

}

func main() {

	var result Interface1
	result = &List{}

	result.Add(10)
	result.Add(12)
	result.Add(13)
	result.Add(14)
	result.Add(15)
	result.Add(15)
	fmt.Println(result)
	result.Update(15, 20)
	fmt.Println(result)
	result.DeleteOne(20)
	fmt.Println(result)
	result.DeleteOne(20)
	fmt.Println(result)
}
