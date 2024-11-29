package main

import "fmt"

func main() {

	var beratIkan [1000]float64

	var jumlahIkan, banyakIkanPerWadah int

	// Input jumlah ikan yang akan dijual
	fmt.Print("Masukkan jumlah ikan yang akan dijual: ")
	fmt.Scanln(&jumlahIkan)

	// Input banyaknya ikan di setiap wadah
	fmt.Print("Masukkan banyaknya ikan di setiap wadah: ")
	fmt.Scanln(&banyakIkanPerWadah)

	fmt.Println("Masukkan berat ikan (masing-masing dipisahkan dengan spasi): ")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Scanf("%f", &beratIkan[i])
	}

	// Hitung total berat ikan di setiap wadah
	var totalBeratPerWadah []float64
	for i := 0; i < banyakIkanPerWadah; i++ {
		var totalBerat float64
		for j := i * banyakIkanPerWadah; j < (i+1)*banyakIkanPerWadah && j < jumlahIkan; j++ {
			totalBerat += beratIkan[j]
		}
		totalBeratPerWadah = append(totalBeratPerWadah, totalBerat)
	}

	// Hitung berat rata-rata ikan di setiap wadah
	var beratRataRata float64
	for _, berat := range totalBeratPerWadah {
		beratRataRata += berat
	}
	beratRataRata /= float64(banyakIkanPerWadah)

	// Output hasil
	fmt.Println("Total berat ikan di setiap wadah:")
	for _, berat := range totalBeratPerWadah {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	fmt.Printf("%.2f", beratRataRata)
}
