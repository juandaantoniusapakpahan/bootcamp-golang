package obat

import "strconv"

type Obat struct {
	Id        string
	NamaObat  string
	HargaObat float32
}

type ObatInterface interface {
	Add(namaobat string, hargaobat float32)
	FindObatById(obatId string) Obat
	FindObatByName(namaobat string) Obat
}

type ListObat struct {
	Obats []Obat
}

func NewObat() ObatInterface {
	return &ListObat{}
}

func (o *ListObat) Add(namaobat string, hargaobat float32) {
	id := "obat-" + strconv.Itoa(len(o.Obats))
	newObat := Obat{Id: id, NamaObat: namaobat, HargaObat: hargaobat}
	o.Obats = append(o.Obats, newObat)
}

func (o *ListObat) FindObatById(obatId string) Obat {
	for _, v := range o.Obats {
		if v.Id == obatId {
			return v
		}
	}
	return Obat{}
}

func (o *ListObat) FindObatByName(namaobat string) Obat {
	for _, v := range o.Obats {
		if v.NamaObat == namaobat {
			return v
		}
	}
	return Obat{}
}
