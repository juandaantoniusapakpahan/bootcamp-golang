package main

import (
	"fmt"
	"go-database-setting/database"
)

func main() {
	student := database.Student{
		FirstName:  "Juanda",
		MiddleName: "Antonius",
		LastName:   "Pakpahan",
	}

	database.Insert(&student, database.MyDB())
	result := database.FindByName(student.FirstName, student.LastName, database.MyDB())
	fmt.Println(result)
	fmt.Println("Completed")
}
