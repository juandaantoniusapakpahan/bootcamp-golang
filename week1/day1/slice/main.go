// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func main() {
	kota := [5]string{"Jakarta", "Bogor", "Depok", "Tanggerang", "Bekasi"}
	// Data Awal
	fmt.Println("Awal Kota:", kota, "\t\t", "Cap:", cap(kota))

	// Soal1
	slice1 := kota[1:]
	slice1 = append(slice1, "Medan")
	fmt.Println("Hasil Soal 1:", slice1, "\t", "Cap:", cap(slice1))

	// Soal2
	slice2 := slice1[:]
	slice2 = append([]string{"Bandung"}, slice2...)
	slice2 = append(slice2[:len(slice2)-2], slice2[len(slice2)-1:]...)
	fmt.Println("Hasil Soal 2:", slice2, "\t", "Cap:", cap(slice2))

	// Soal3
	slice3 := slice2[:]
	slice3[len(slice2)-1] = "Semarang"
	fmt.Println("Hasil Soal 3:", slice3, "", "Cap:", cap(slice3))

}
