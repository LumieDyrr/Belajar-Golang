package main

import "fmt"

func main() {

	//sturct
	pelanggan1 := NewKeranjang("Dafa")
	items := Menu{
		id:      1,
		makanan: "Soto",
		minuman: "Thai Tea",
		jumlah:  2,
		harga:   30000,
	}
	pelanggan1.InsertMenu(items)
	items2 := []Menu{
		{
			id:      2,
			makanan: "Nasi Goreng",
			minuman: "Es Teh",
			jumlah:  3,
			harga:   35000,
		},
		{
			id:      3,
			makanan: "Mie Goreng",
			minuman: "Teh Tarik",
			jumlah:  2,
			harga:   20000,
		},
	}
	pelanggan1.InsertMenuvar(items2...)

	update := Menu{
		id:      2,
		makanan: "Nasi Goreng",
		minuman: "Boba Ice",
		jumlah:  2,
		harga:   30000,
	}
	pelanggan1.UpdateMenu(update)
	pelanggan1.DeleteMenu(1)
	fmt.Println(pelanggan1)

	//Interface
	Kt := KaryawanTetap{"Dafa", 4500000, 250000}
	Kk := KaryawanKontrak{"Aisyah", 4000000}
	free := Freelance{"Almas", 300000, 10}
	fmt.Println(Kt.Gaji())
	fmt.Println(Kk.Gaji())
	fmt.Println(free.Gaji())

	Total := Perusahaan(&Kk, &Kt, &free)
	fmt.Println(Total)
}
