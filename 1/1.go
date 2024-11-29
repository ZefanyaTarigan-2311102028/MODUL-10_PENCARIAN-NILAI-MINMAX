package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan jumlah anak kelinci
	var n int
	// Deklarasi variabel untuk menyimpan berat anak kelinci
	var berat []float64

	// Meminta input jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scanln(&n)

	// Meminta input berat anak kelinci
	berat = make([]float64, n)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&berat[i])
	}

	// Mencari berat kelinci terkecil dan terbesar
	terkecil := berat[0]
	terbesar := berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < terkecil {
			terkecil = berat[i]
		}
		if berat[i] > terbesar {
			terbesar = berat[i]
		}
	}

	// Menampilkan hasil
	fmt.Println("Berat kelinci terkecil:", terkecil)
	fmt.Println("Berat kelinci terbesar:", terbesar)
}
