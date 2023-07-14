package penanganan

import (
	"fmt"
	"hospital-exam/dokter"
	"hospital-exam/pasien"
	"hospital-exam/penyakit"
)

// Penanganan → Pasien pasien, Penyakit []penyakit, Dokter dokter, HargaPelayanan float

type Penanganan struct {
	Pasien         pasien.Pasien
	Penyakit       []penyakit.Penyakit
	Dokter         dokter.Dokter
	HargaPelayanan float32
}

type PenangananInterfance interface {
	Add(namapasien string, namapenyakit []string, namadokter string, hargapelayanan float32)
	ShowAll()
}

type ListPenanganan struct {
	Penanganans []Penanganan
	Pasien      pasien.PasienInterfance
	Penyakit    penyakit.PenyakitInterface
	Dokter      dokter.DokterInterfance
}

func NewPenanganan(pasien pasien.PasienInterfance, penyakit penyakit.PenyakitInterface, dokter dokter.DokterInterfance) PenangananInterfance {
	return &ListPenanganan{Pasien: pasien, Penyakit: penyakit, Dokter: dokter}
}

func (p *ListPenanganan) ShowAll() {
	for _, v := range p.Penanganans {
		fmt.Println(v)
	}
}

func (p *ListPenanganan) Add(namapasien string, namapenyakit []string, namadokter string, hargapelayanan float32) {
	dataPasien := p.Pasien.FindPasienByName(namapasien)
	if dataPasien == (pasien.Pasien{}) {
		fmt.Println("Pasien tidak ditemukan")
		return
	}

	dataPenyakit := p.Penyakit.FindPenyakitByName(namapenyakit)
	if len(dataPenyakit) == 0 {
		fmt.Println("Penyakit tidak ditemukan di database")
		return
	}

	dataDokter := p.Dokter.FindDokterByNama(namadokter)
	if dataDokter == (dokter.Dokter{}) {
		fmt.Println("Dokter tidak ditemukan di database")
		return
	}

	newPenanganan := Penanganan{
		Pasien:         dataPasien,
		Penyakit:       dataPenyakit,
		Dokter:         dataDokter,
		HargaPelayanan: hargapelayanan,
	}

	p.Penanganans = append(p.Penanganans, newPenanganan)
}
