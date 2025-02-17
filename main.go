package main

import "fmt"

func main() {

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
