package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type DBTest struct {
}

type DBInterface interface {
	CheckFile(filename string) error
	MarshalMan(filename string, data interface{})
	ReadFile(filename string, data interface{})
	GetFile(filename string) *os.File
	OpenLogFile(filename string)
}

func NewDB() DBInterface {
	return &DBTest{}
}

func (d *DBTest) CheckFile(filename string) error {
	_, err := os.Stat("db/" + filename)
	if os.IsNotExist(err) {
		_, err := os.Create("db/" + filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DBTest) MarshalMan(filename string, data interface{}) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("db/"+filename, dataBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (d *DBTest) ReadFile(filename string, data interface{}) {
	if err := d.CheckFile(filename); err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile("db/" + filename)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, data)
}

func (d *DBTest) GetFile(filename string) *os.File {
	file, err := os.Open("db/" + filename)
	if err != nil {
		panic(err)
	}
	return file
}

func (d *DBTest) OpenLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}
