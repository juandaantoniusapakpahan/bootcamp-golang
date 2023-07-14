package dokter

import "strconv"

type Dokter struct {
	Id              string
	NamaDokter      string
	Spesialisasi    string
	TarifKonsultasi float32
}

type DokterInterfance interface {
	Add(namadokter string, spesialisasi string, tarifkonsultasi float32)
	FindDokterById(dokterId string) Dokter
	FindDokterByNama(namadokter string) Dokter
}

type ListDokter struct {
	Dokters []Dokter
}

func NewDokter() DokterInterfance {
	return &ListDokter{}
}

func (d *ListDokter) FindDokterByNama(namadokter string) Dokter {
	for _, v := range d.Dokters {
		if v.NamaDokter == namadokter {
			return v
		}
	}
	return Dokter{}
}
func (d *ListDokter) FindDokterById(dokterId string) Dokter {

	for _, v := range d.Dokters {
		if v.Id == dokterId {
			return v
		}
	}
	return Dokter{}
}

func (d *ListDokter) Add(namadokter string, spesialisasi string, tarifkonsultasi float32) {
	id := "dokter-" + strconv.Itoa(len(d.Dokters))
	newDokter := Dokter{Id: id, NamaDokter: namadokter, Spesialisasi: spesialisasi, TarifKonsultasi: tarifkonsultasi}
	d.Dokters = append(d.Dokters, newDokter)
}
