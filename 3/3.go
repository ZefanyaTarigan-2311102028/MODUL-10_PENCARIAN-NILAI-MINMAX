package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < len(arrBerat); i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
	fmt.Printf("Berat balita minimum: %.2f kg\n", *bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", *bMax)
}

func rerata(arrBerat arrBalita) float64 {
	var sum float64
	for i := 0; i < len(arrBerat); i++ {
		sum += arrBerat[i]
	}
	return sum / float64(len(arrBerat))
}

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&n)

	var arrBalita arrBalita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&arrBalita[i])
	}

	var bMin, bMax float64
	hitungMinMax(arrBalita, &bMin, &bMax)

	rerataBalita := rerata(arrBalita)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBalita)
}
