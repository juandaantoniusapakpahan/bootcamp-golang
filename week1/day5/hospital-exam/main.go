package main

import (
	"hospital-exam/dokter"
	"hospital-exam/obat"
	"hospital-exam/pasien"
	"hospital-exam/penanganan"
	"hospital-exam/penyakit"
)

func main() {
	newObat := obat.NewObat()
	newPenyakit := penyakit.NewPenyakit(newObat)
	newPasien := pasien.NewPasien()
	newDokter := dokter.NewDokter()
	newPenanganan := penanganan.NewPenanganan(newPasien, newPenyakit, newDokter)

	newObat.Add("AAAA", 30000)
	newObat.Add("BBBB", 40000)
	newObat.Add("CCCC", 50000)

	newPenyakit.Add("DDDD", "AAAA")
	newPenyakit.Add("DDDD", "BBBB")

	newDokter.Add("Dokter A", "Bagian Dalam", 200000)
	newDokter.Add("Dokter B", "THT", 200000)

	newPasien.Add("Rojo", "29", "Jalan Pelita Harapan", "DDDD")

	newPenanganan.Add("Rojo", []string{"DDDD"}, "Dokter A", 50000)

	newPenanganan.ShowAll()

}
