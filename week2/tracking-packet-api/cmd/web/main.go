package main

import (
	"trancking-packet/pkg/db"
	"trancking-packet/pkg/ihttp/handler"
	"trancking-packet/pkg/ihttp/router"

	"github.com/go-playground/validator/v10"
)

func main() {
	newDb := db.NewDB()
	newValidate := validator.New()
	newPengirim := db.NewPengirim()
	newPenerima := db.NewPenerima()
	newService := db.NewService()
	newLokasi := db.NewLokasi()
	penerimaHandler := handler.NewPenerimaHandler(newPenerima, newDb, newValidate)
	pengirimHandler := handler.NewPengirimHandler(newPengirim, newDb, newValidate)
	serviceHandler := handler.NewServiceHandler(newService, newDb, newValidate)
	lokasiHandler := handler.NewLokasiHandler(newLokasi, newDb, newValidate)

	route := router.NewRoute(penerimaHandler, pengirimHandler, serviceHandler, lokasiHandler)
	route.Run()

}
