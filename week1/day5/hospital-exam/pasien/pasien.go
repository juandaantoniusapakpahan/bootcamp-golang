package pasien

import "strconv"

// Pasien → IdPasien string, NamaPasien string,
// UmurPasien string, AlamatPasien string, Penyakit string

type Pasien struct {
	Id           string
	NamaPasien   string
	UmurPasien   string
	AlamatPasien string
	Penyakit     string
}

type PasienInterfance interface {
	Add(namapasien string, umurpasien string, alamanpasien string, penyakit string)
	FindPasienByName(namapasien string) Pasien
	FindPasienById(pasienId string) Pasien
}

type ListPasien struct {
	Pasiens []Pasien
}

func NewPasien() PasienInterfance {
	return &ListPasien{}
}

func (p *ListPasien) FindPasienById(pasienId string) Pasien {
	for _, v := range p.Pasiens {
		if v.Id == pasienId {
			return v
		}
	}
	return Pasien{}
}

func (p *ListPasien) FindPasienByName(namapasien string) Pasien {
	for _, v := range p.Pasiens {
		if v.NamaPasien == namapasien {
			return v
		}
	}
	return Pasien{}
}

func (p *ListPasien) Add(namapasien string, umurpasien string, alamanpasien string, penyakit string) {
	id := "pasien-" + strconv.Itoa(len(p.Pasiens))
	newPasien := Pasien{
		Id:           id,
		NamaPasien:   namapasien,
		UmurPasien:   umurpasien,
		AlamatPasien: alamanpasien,
		Penyakit:     penyakit,
	}

	p.Pasiens = append(p.Pasiens, newPasien)
}
