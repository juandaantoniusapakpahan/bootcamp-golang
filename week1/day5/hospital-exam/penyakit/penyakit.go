package penyakit

import (
	"fmt"
	"hospital-exam/obat"
	"strconv"
)

type Penyakit struct {
	Id           string
	NamaPenyakit string
	Obat         obat.Obat
}

type PenyakitInterface interface {
	Add(namapenyakit string, namaobat string)
	FindPenyakitById(penyakitId string) Penyakit
	FindPenyakitByName(namapenyakit []string) []Penyakit
}

type ListPenyakitImpl struct {
	Penyakits []Penyakit
	Obat      obat.ObatInterface
}

func NewPenyakit(obat obat.ObatInterface) PenyakitInterface {
	return &ListPenyakitImpl{Obat: obat}
}

func (p *ListPenyakitImpl) FindPenyakitByName(namapenyakit []string) []Penyakit {
	lenNamaPenyakit := 0
	lenPenyakit := 0
	penyakits := []Penyakit{}

	for lenNamaPenyakit < len(namapenyakit) {
		if namapenyakit[lenNamaPenyakit] == p.Penyakits[lenPenyakit].NamaPenyakit {
			penyakits = append(penyakits, p.Penyakits[lenPenyakit])
			lenPenyakit = 0
			lenNamaPenyakit++
			continue
		}
		lenPenyakit++
	}

	return penyakits
}

func (p *ListPenyakitImpl) FindPenyakitById(penyakitId string) Penyakit {
	for _, v := range p.Penyakits {
		if v.Id == penyakitId {
			return v
		}
	}
	return Penyakit{}
}

func (p *ListPenyakitImpl) Add(namapenyakit string, namaobat string) {
	dataobat := p.Obat.FindObatByName(namaobat)

	if dataobat == (obat.Obat{}) {
		fmt.Println("Tidak ada Obat nya Gan")
		return
	}

	id := "penyakit-" + strconv.Itoa(len(p.Penyakits))
	newPenyakit := Penyakit{Id: id, NamaPenyakit: namapenyakit, Obat: dataobat}

	p.Penyakits = append(p.Penyakits, newPenyakit)
}
