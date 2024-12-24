package main

import "fmt"

type Pabrikan struct {
	ID          int
	Nama        string
	Wilayah     string
	HitungMobil int
	TotalSales  int
}

type Mobil struct {
	ID_mobil int
	NamaMobil string
	IDpabrikan int
	Tahun int
	HargaMobil float64
}

const maxCars = 100
var mobilList [maxCars]Mobil

const maxPabrikan = 100 // Maksimum jumlah pabrikan
var pabrikanList [maxPabrikan]Pabrikan
var pabrikanCount int = 0 // Jumlah pabrikan saat ini

var mobilCount int = 0
var nextCarID int = 1

// Fungsi Untuk Menambahkan Pabrikan
func tambahPabrik(nama, negara string) {
	if pabrikanCount >= maxPabrikan {
		fmt.Println("Tidak dapat menambahkan pabrikan baru. Kapasitas penuh!")
		return
	}
	Manufaktur := Pabrikan{
		ID: pabrikanCount + 1,
		Nama: nama,
		Wilayah: negara,
		HitungMobil: 0,
	}
	pabrikanList[pabrikanCount] = Manufaktur
	pabrikanCount++
	fmt.Println("Pabrikan berhasil ditambahkan!")
}

// Fungsi untuk mengubah data pabrikan
func ubahPabrikan(id int, nama, negara string) {
	for i := 0; i < pabrikanCount; i++ {
		if pabrikanList[i].ID == id {
			pabrikanList[i].Nama = nama
			pabrikanList[i].Wilayah = negara
			fmt.Println("Pabrikan berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan!")
}

// Fungsi menghapus pabrikan dan mobilnya.
func hapusPabrikan(id int) {
	// Hapus pabrikan
	for i := 0; i < pabrikanCount; i++ {
		if pabrikanList[i].ID == id {
			// Geser elemen-elemen setelah yang dihapus
			for j := i; j < pabrikanCount-1; j++ {
				pabrikanList[j] = pabrikanList[j+1]
			}
			pabrikanCount-- // Kurangi jumlah elemen
			fmt.Println("Pabrikan ditemukan dan dihapus.")
			break
		}
	}

	// Hapus mobil terkait
	for i := 0; i < mobilCount; {
		if mobilList[i].IDpabrikan == id {
			// Geser elemen-elemen setelah yang dihapus
			for j := i; j < mobilCount-1; j++ {
				mobilList[j] = mobilList[j+1]
			}
			mobilCount-- // Kurangi jumlah elemen
			fmt.Println("Mobil terkait berhasil dihapus.")
		} else {
			i++ // Hanya lanjut jika elemen tidak dihapus
		}
	}
}


// Fungsi Untuk Menambahkan Mobil Ke Pabrikan
func tambahMobil(namaMobil string, idPabrikan, tahun int, harga float64) {
	// Cari pabrikan berdasarkan ID
	for i := 0; i < pabrikanCount; i++ {
		if pabrikanList[i].ID == idPabrikan {
			if mobilCount >= maxCars {
				fmt.Println("Tidak dapat menambahkan mobil baru. Kapasitas penuh!")
				return
			}

			// Tambahkan mobil ke array
			mobilList[mobilCount] = Mobil{
				ID_mobil: nextCarID,
				NamaMobil: namaMobil,
				IDpabrikan: idPabrikan,
				Tahun: tahun,
				HargaMobil: harga,
			}
			mobilCount++
			pabrikanList[i].HitungMobil++
			nextCarID++
			fmt.Println("Mobil berhasil ditambahkan!")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan!")
}

// ubahMobil mengedit data mobil yang ada.
func ubahMobil(id int, nama string, idPabrikan, tahun int, harga float64) {
	for i := 0; i < mobilCount; i++ {
		if mobilList[i].ID_mobil == id {
			// Perbarui data mobil
			mobilList[i].NamaMobil = nama
			mobilList[i].IDpabrikan = idPabrikan
			mobilList[i].Tahun = tahun
			mobilList[i].HargaMobil = harga
			fmt.Println("Mobil berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Mobil tidak ditemukan!")
}

// Funsgi Untuk Menampilkan Data Mobil
func tampilkanMobil() {
	for i := 0; i < mobilCount; i++ {
		fmt.Printf("\nID Mobil : %d Nama: %s ID Pabrikan: %d\nTahun: %d\nHarga: %.2f\n",
			mobilList[i].ID_mobil, mobilList[i].NamaMobil, mobilList[i].IDpabrikan, mobilList[i].Tahun, mobilList[i].HargaMobil)
		fmt.Println()
	}
}

// Fungsi Menghapus Data Mobil
func hapusMobil(id int) {
	for i := 0; i < mobilCount; i++ {
		if mobilList[i].IDpabrikan == id {
			// Geser elemen-elemen setelah elemen yang dihapus
			for j := i; j < mobilCount-1; j++ {
				mobilList[j] = mobilList[j+1]
			}
			mobilCount-- // Kurangi jumlah mobil
			fmt.Println("Mobil berhasil dihapus!")
			return
		}
	}
	fmt.Println("Mobil tidak ditemukan!")
}

// Fungsi Untuk Menampilkan Pabrikan
func tampilkanPabrikan() {
	for i := 0; i < pabrikanCount; i++ {
		fmt.Printf("ID: %d\n Nama: %s\n Wilayah: %s\n Hitung Mobil: %d\n",
			pabrikanList[i].ID, pabrikanList[i].Nama, pabrikanList[i].Wilayah, pabrikanList[i].HitungMobil)
			fmt.Println()
	}
	
}

func main() {
for {
	fmt.Println("1.Tambah Pabrikan")
	fmt.Println("2.Ubah Data Pabrikan")
	fmt.Println("3.Tampilkan Pabrikan")
	fmt.Println("4.Hapus Pabrikan")
	fmt.Println("5.Tambah Mobil")
	fmt.Println("6.Ubah Data Mobil")
	fmt.Println("7.Tampilkan Mobil")
	fmt.Println("8.Hapus Mobil")
	fmt.Print("Pilih Nomor : ")
	var number int
	fmt.Scan(&number)
	switch number {
	case 1 :
		var nama, negara string
		fmt.Print("\nInput nama Pabrikan : ")
		fmt.Scan(&nama)
		fmt.Print("Input Wilayah Pabrikan : ")
		fmt.Scan(&negara)
		tambahPabrik(nama, negara)
		fmt.Println()

	case 2 :
		fmt.Println()
		var id int
		var nama, negara string
		fmt.Print("Masukan ID Pabrikan Yang Akan Di ubah : ")
		fmt.Scan(&id)
		fmt.Print("Nama Pabrikan    : ")
		fmt.Scan(&nama)
		fmt.Print("Wilayah Pabrikan : ")
		fmt.Scan(&negara)
		ubahPabrikan(id, nama, negara)
		fmt.Println()

	case 3 :
		fmt.Println()
		tampilkanPabrikan()

	case 4 : 
		fmt.Println()
		var id int
		fmt.Print("Input ID Pabrikan Yang Akan Dihapus : ")
		fmt.Scan(&id)
		hapusPabrikan(id)
		fmt.Println()

	case 5 :
		var namaMobil string
		var idPabrikan, tahun int
		var harga_mobil float64
		fmt.Print("\nNama Mobil : ")
		fmt.Scan(&namaMobil)
		fmt.Print("Tambahkan dipabrikan ke- : ")
		fmt.Scan(&idPabrikan)
		fmt.Print("Tahun : ")
		fmt.Scan(&tahun)
		fmt.Print("Harga Mobil : Rp ")
		fmt.Scan(&harga_mobil)
		tambahMobil(namaMobil, idPabrikan, tahun, harga_mobil)
		fmt.Println()

	case 6 : 
		var id, idPabrikan, tahun int
		var nama string
		var harga float64
		fmt.Print("Input ID mobil yang akan di ubah : ")
		fmt.Scan(&id)
		fmt.Print("Ubah Nama Mobil : ")
		fmt.Scan(&nama)
		fmt.Print("Ubah ID Pabrikan : ")
		fmt.Scan(&idPabrikan)
		fmt.Print("Ubah Tahun : ")
		fmt.Scan(&tahun)
		fmt.Print("Ubah Harga : ")
		fmt.Scan(&harga)
		ubahMobil(id, nama, idPabrikan, tahun, harga)
		fmt.Println()

	case 7 : 
		fmt.Println()
		tampilkanMobil()
		fmt.Println()
		
	case 8 : 
		var id_mobil int
		fmt.Print("Inputkan ID Mobil Untuk Menghapus : ")
		fmt.Scan(&id_mobil)
		hapusMobil(id_mobil)
		
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
  }
}
