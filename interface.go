package main

type Karyawan interface {
	Gaji() int
}

type KaryawanTetap struct {
	Nama      string
	GajiPokok int
	Tunjangan int
}

type KaryawanKontrak struct {
	Nama      string
	GajiPokok int
}

type Freelance struct {
	Nama       string
	GajiPerjam int
	TotalJam   int
}

func (kt *KaryawanTetap) Gaji() int {
	return kt.GajiPokok + kt.Tunjangan
}

func (Kk *KaryawanKontrak) Gaji() int {
	return Kk.GajiPokok
}

func (F *Freelance) Gaji() int {
	return F.GajiPerjam * F.TotalJam
}

func Perusahaan(gaji ...Karyawan) int {
	total := 0
	for _, value := range gaji {
		total += value.Gaji()
	}
	return total
}
