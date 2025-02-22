package main

import "fmt"

type Penjualan struct {
	IDTransaksi, JumlahPenjualan int
}

func shellSort(data []Penjualan) {
	n := len(data)

	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := data[i]
			j := i
			for j >= gap && data[j-gap].JumlahPenjualan < temp.JumlahPenjualan {
				data[j] = data[j-gap]
				j -= gap
			}
			data[j] = temp
		}
	}
}

func main() {
	var penjualan = []Penjualan{
		{1, 150},
		{2, 200},
		{3, 50},
		{4, 300},
		{5, 100},
		{6, 450},
	}

	fmt.Println("Data Penjualan Sebelum Diurutkan:")
	for _, p := range penjualan {
		fmt.Printf("ID Transaksi: %d, Jumlah Penjualan: %d\n", p.IDTransaksi, p.JumlahPenjualan)
	}

	shellSort(penjualan)

	fmt.Println("\nData Penjualan Setelah Diurutkan:")
	for _, p := range penjualan {
		fmt.Printf("ID Transaksi: %d, Jumlah Penjualan: %d\n", p.IDTransaksi, p.JumlahPenjualan)
	}
}

/*
func main() {
	names := []string{"almas", "dafa", "azril", "Aisyah"}

	for indeks, value := range names {
		var result = indeks
		var values = value
		if values == "Aisyah" {
			fmt.Printf("%v berada pada indeks ke %v\n", values, result)
			break
		}
		//fmt.Printf("the value at pos %v  is %v\n", result, values)
	}
}
*/
