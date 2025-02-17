package main

type Pelanggan struct {
	Name      string
	keranjang []Menu
	total     float64
}

type Menu struct {
	id      int
	makanan string
	minuman string
	jumlah  int
	harga   float64
}

func NewKeranjang(nama string) Pelanggan {
	return Pelanggan{
		Name:      nama,
		keranjang: []Menu{},
		total:     0,
	}
}

func (pelanggan *Pelanggan) InsertMenu(menu Menu) {
	pelanggan.keranjang = append(pelanggan.keranjang, menu)
}

func (pelanggan *Pelanggan) InsertMenuvar(menu ...Menu) {
	pelanggan.keranjang = append(pelanggan.keranjang, menu...)
}

func (pelanggan *Pelanggan) UpdateMenu(menu Menu) {
	for indeks, value := range pelanggan.keranjang {
		if value.id == menu.id {
			pelanggan.keranjang[indeks] = menu
		}
	}
	pelanggan.TotalMenu()
}

func (pelanggan *Pelanggan) DeleteMenu(id int) {
	for indeks, value := range pelanggan.keranjang {
		if value.id == id {
			pelanggan.keranjang = append(pelanggan.keranjang[:indeks], pelanggan.keranjang[indeks+1:]...)
		}
	}
	pelanggan.TotalMenu()
}

func (pelanggan *Pelanggan) TotalMenu() {
	var total float64
	for _, value := range pelanggan.keranjang {
		total += value.harga * float64(value.jumlah)
	}
	pelanggan.total = total
}
