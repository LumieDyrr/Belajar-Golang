package main

import "fmt"

func main() {

	//Struct
	result := NewKeranjang("Dafa")
	menu := Menu{
		id:      1,
		makanan: "Nasi goreng",
		minuman: "Es Teh",
		jumlah:  2,
		harga:   25000,
	}
	result.InsertMenu(menu)

	tambah := []Menu{
		{
			id:      2,
			makanan: "Sate",
			minuman: "Bir",
			jumlah:  2,
			harga:   35000,
		},
		{
			id:      3,
			makanan: "Soto",
			minuman: "Kopi",
			jumlah:  2,
			harga:   15000,
		},
	}

	result.InsertMenuvar(tambah...)
	update := Menu{
		id:      2,
		makanan: "Coto",
		minuman: "kopi",
		jumlah:  2,
		harga:   45000,
	}
	result.UpdateMenu(update)
	result.DeleteMenu(1)
	fmt.Println(result)

}
