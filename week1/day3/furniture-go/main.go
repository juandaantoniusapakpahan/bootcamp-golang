package main

import (
	"errors"
	"fmt"
)

type Barang struct {
	Id    int64
	Nama  string
	Merek string
	Harga float32
	Stok  int
	Next  *Barang
}

type ListKatalog struct {
	Data  *Barang
	Total int
}

type InterfaceI interface {
	Print()
	Add(data Barang)
	RemoveByName(name string) error
	Search(id int64) (*Barang, error)
	Update(id int64, barang Barang) error
}

func (k *ListKatalog) Print() {
	if k.Data == nil {
		return
	}

	cur := k.Data
	type M map[string]interface{}
	for cur != nil {

		fmt.Println(M{"id": cur.Id, "nama": cur.Nama, "merek": cur.Merek, "harga": cur.Harga, "stok": cur.Stok})

		cur = cur.Next
	}
}

func (l *ListKatalog) Add(data Barang) {
	l.Total = l.Total + 1
	data.Id = int64(l.Total)

	if l.Data == nil {

		l.Data = &data
		return
	}

	cur := l.Data
	for cur.Next != nil {
		cur = cur.Next
	}

	if cur.Next == nil {
		cur.Next = &data
	}
}

func (l *ListKatalog) RemoveByName(name string) error {
	if l.Data == nil {
		return errors.New("Barang tidak ditemukan")
	}

	if l.Data.Nama == name {
		l.Data = l.Data.Next
		return nil
	}

	cur := l.Data
	for cur.Next != nil && cur.Next.Nama != name {
		cur = cur.Next
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	return errors.New("Barang tidak ditemukan")
}

func (l *ListKatalog) Search(id int64) (*Barang, error) {
	if l.Data == nil {
		return &Barang{}, errors.New("Barang tidak ditemukan")
	}

	if l.Data.Id == id {
		return l.Data, nil
	}

	cur := l.Data
	for cur.Next != nil && cur.Next.Id != id {
		cur = cur.Next
	}
	barang := &Barang{}
	if cur.Next != nil {
		barang = cur.Next
	}

	return barang, nil
}

func (l *ListKatalog) Update(id int64, barang Barang) error {
	data, err := l.Search(id)
	if err != nil {
		return err
	}

	data.Nama = barang.Nama
	data.Merek = barang.Merek
	data.Harga = barang.Harga
	data.Stok = barang.Stok

	return nil
}

func main() {

	list := ListKatalog{}
	var inter InterfaceI = &list

	kata1 := Barang{
		Nama:  "A",
		Merek: "AA",
		Harga: 12,
		Stok:  10,
	}
	kata2 := Barang{
		Nama:  "B",
		Merek: "BA",
		Harga: 12,
		Stok:  10,
	}
	kata3 := Barang{
		Nama:  "C",
		Merek: "CA",
		Harga: 12,
		Stok:  10,
	}
	kata4 := Barang{
		Nama:  "D",
		Merek: "DA",
		Harga: 12,
		Stok:  10,
	}
	kataUpdate := Barang{
		Nama:  "D",
		Merek: "DA",
		Harga: 12,
		Stok:  10,
	}
	inter.Add(kata1)
	inter.Add(kata2)
	inter.Add(kata3)
	inter.Add(kata4)

	inter.Print()
	fmt.Println()

	err := inter.RemoveByName("A")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Barang berhasil dihapus")
	}
	inter.Print()

	fmt.Println()
	fmt.Println("Search barang")

	barang, err := inter.Search(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*barang)
	}
	fmt.Println()

	err = inter.Update(2, kataUpdate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Update success")
	}
	fmt.Println()
	inter.Print()
}
