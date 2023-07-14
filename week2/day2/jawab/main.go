package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type MetaData struct {
	Sku string `json:"sku"`
}

type BankAccount struct {
	BankAccountHash         string `json:"bank_account_hash"`
	MaskedBankAccountNumber string `json:"masked_bank_account_number"`
}

type ChannelProperties struct {
	FailureReturnUrl string `json:"failure_return_url"`
	SuccessReturnUrl string `json:"success_return_url"`
}

type DirectDebit struct {
	Type              string            `json:"type"`
	DebitCard         string            `json:"debit_card"`
	BankAccount       BankAccount       `json:"bank_account"`
	ChannelCode       string            `json:"channel_code"`
	ChannelProperties ChannelProperties `json:"channel_properties"`
}

type PaymentMethod struct {
	Id                 string      `json:"id"`
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	Created            time.Time   `json:"created"`
	Ewallet            string      `json:"ewallet,omitempty"`
	QrCode             string      `json:"qr_code,omitempty"`
	Updated            string      `json:"updated"`
	Description        string      `json:"description"`
	Reusability        string      `json:"reusability"`
	DirectDebit        DirectDebit `json:"direct_debit"`
	ReferenceId        string      `json:"reference_id"`
	VirtualAccount     string      `json:"virtual_account"`
	OverTheCounter     string      `json:"over_the_counter"`
	DirectBankTransfer string      `json:"direct_bank_transfer"`
}

type Person struct {
	Id                string        `json:"id"`
	Amount            int           `json:"amount"`
	Country           string        `json:"country"`
	Created           time.Time     `json:"created"`
	Updated           string        `json:"updated"`
	Currency          string        `json:"currency"`
	MetaData          MetaData      `json:"metadata"`
	CustomerId        string        `json:"customer_id"`
	ReferenceId       string        `json:"reference_id"`
	PaymentMethod     PaymentMethod `json:"payment_method"`
	Description       string        `json:"description,omitempty"`
	FailureCode       string        `json:"failure_code"`
	PaymentDetail     string        `json:"payment_detail"`
	ChannelProperties string        `json:"channel_properties"`
	PaymentRequestId  string        `json:"payment_request_id"`
}

func ReadFile() *os.File {
	write, err := os.Open("data.json")

	if err != nil {
		panic(err)
	}
	return write
}

func NewRead() []byte {
	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	return content
}

func Decode() Person {

	write := ReadFile()
	defer write.Close()
	result := Person{}
	decode := json.NewDecoder(write)
	decode.Decode(&result)

	return result
}

func Unmarshal() Person {

	result := Person{}
	ggwp := NewRead()
	err := json.Unmarshal(ggwp, &result)
	if err != nil {
		panic(err)
	}
	return result
}

type M map[string]interface{}

func UnmarshalMap() M {
	result := M{}
	ggwp := NewRead()
	err := json.Unmarshal(ggwp, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func NewDecodeMap() M {

	write := ReadFile()
	result := make(M)
	decode := json.NewDecoder(write)
	decode.Decode(&result)

	return result
}

// Buatlah struct Book yang berisi field Title string, ISBN []string,
// Author string, PublishDate string.
// Kemudian buatlah menjadi jadi bentuk JSON menggunakan:

type Book struct {
	Title       string    `json:"title"`
	ISBN        []string  `json:"isbn"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publishDate"`
}

func MarshalGo() string {
	book := Book{
		Title:       "GGWP",
		ISBN:        []string{"js9fuas98df", "oiasdfjs9"},
		Author:      "Smith",
		PublishDate: time.Now(),
	}

	byte, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	result := string(byte)
	return result
}

func NewEncoder() string {
	book := Book{
		Title:       "GGWP",
		ISBN:        []string{"js9fuas98df", "oiasdfjs9"},
		Author:      "Smith",
		PublishDate: time.Now(),
	}
	write := new(bytes.Buffer)
	encoder := json.NewEncoder(write)
	encoder.Encode(book)
	return write.String()
}

func main() {
	nomor1 := Decode()
	nomor2 := Unmarshal()
	nomor3 := UnmarshalMap()
	nomor4 := NewDecodeMap()
	nomor5 := MarshalGo()
	nomor6 := NewEncoder()

	fmt.Println(nomor1.Created)
	fmt.Println()
	fmt.Println(nomor2)
	fmt.Println()
	fmt.Println(nomor3)
	fmt.Println()
	fmt.Println(nomor4)
	fmt.Println()
	fmt.Println(nomor5)
	fmt.Println()
	fmt.Println(nomor6)
	fmt.Println()

	var gg string = `{
		"person": 
		[
			{
				"name":"ggwp"
			},
			{
				"name":"cokai"
			}
		]
	}`

	type PersonI struct {
		Name string `json:"name"`
	}
	type PersonII struct {
		Person []PersonI `json:"person"`
	}

	test := PersonII{}
	json.Unmarshal([]byte(gg), &test)
	fmt.Println(test)
}
