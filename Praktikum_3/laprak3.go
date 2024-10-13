// Soal Latihan Modul 2B
// Nomor 3
// package main
// import ("fmt")
// func selisih(a, b float64) float64{
// 	if a > b {
// 		return a - b
// 	}
// 	return b - a
// }
// func main (){
// 	var berat1, berat2 float64
// 	for berat1 > 0 || berat2 > 0 || berat1 + berat2 != 150 {
// 		fmt.Print("Masukan Berat Belanjaan Di Kedua Kantong : ")
// 		fmt.Scanln(&berat1, &berat2)

// 		if selisih(berat1,berat2) < 9 {
// 		fmt. Println("Sepeda Pak Andi Akan Oleng : false \n")
// 		} else if selisih(berat1,berat2) > 9 {
// 			fmt.Println("Sepeda Pak Andi Akan Oleng : true \n")
// 		}
// 		if berat1 + berat2 > 150{
// 			fmt.Println("Program Selesai...")
// 			break
// 		}
// 	}
// }

// Nomor 4
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func countF(k int) float64 {
// 	Pembilang := math.Pow(float64(4*k+2), 2)
// 	Penyebut := float64((4*k + 1) * (4*k + 3))
// 	return Pembilang / Penyebut
// }

// func countAkarDua(k int) float64 {
// 	var akarDua float64 = 1
// 	for a := 0; a <= k; a++ {
// 		akarDua *= countF(a)
// 	}
// 	return akarDua
// }

// func main() {
// 	var k int
// 	fmt.Print("Nilai k : ")
// 	fmt.Scan(&k)

// 	fmt.Printf("Nilai f(%d): %f\n", k, countF(k))

// 	akarloro := countAkarDua(k)
// 	fmt.Println("Nilai Akar 2 : ", akarloro)
// }

// Soal Latihan Modul 2C
// Nomor 1
// package main
// import ("fmt")

// func main(){
// 	// Deklarasi Variabel
// 	var berat int

// 	// Inputan Berat Dari User
// 	fmt.Print("Berat Parcel : ")
// 	fmt.Scanf("%d",&berat)

// 	beratdalamkilo := berat / 1000 // Berat Dalam Kg
// 	biayajasa := beratdalamkilo * 10000 // Jasa Pengiriman Berdasarkan Berat Kg
// 	beratdalamgram := berat % 1000 // Sisa Berat Dalam Gram
// 	fmt.Printf("Detail Berat : %d Kg + %d gram ", beratdalamkilo, beratdalamgram)

// 	// Kondisi Bila Sisa Berat Lebih Dari 500, Maka Dikenakan biaya Rp.5 per-gram
// 	if beratdalamgram >= 500 {
// 		beratdalamgram_jika_lebih_dari_500g := beratdalamgram * 5
// 		fmt.Printf("\nDetail Biaya : Rp.%d + Rp.%d", biayajasa, beratdalamgram_jika_lebih_dari_500g)
// 		totalbiaya_jika_lebih_dari500g := biayajasa + beratdalamgram_jika_lebih_dari_500g
// 		fmt.Printf("\nTotal Biaya : Rp.%d", totalbiaya_jika_lebih_dari500g )

// 	// Kondisi Bila Sisa Berat Lebih Dari 500, Maka Dikenakan biaya Rp.15 per-gram
// 	} else if beratdalamgram <= 500 {
// 		beratdalamgram_jika_kurang_dari_500g := beratdalamgram * 15
// 		fmt.Printf("\nDetail Biaya : Rp.%d + Rp.%d", biayajasa, beratdalamgram_jika_kurang_dari_500g)
// 		totalbiaya_jika_kurang_dari500g := biayajasa + beratdalamgram_jika_kurang_dari_500g
// 		fmt.Printf("\nTotal Biaya  : Rp.%d", totalbiaya_jika_kurang_dari500g )
// 	}

// }
// Nomor 2
// package main
// import ("fmt")
// func main (){
// 	var nam float64
// 	var nmk string
// 	fmt.Print("Nilai Akhir Mata Kuliah : ")
// 	fmt.Scanln(&nam)

// 	if nam > 80 {
// 		nmk = "A"
// 	}else if nam > 72.5 {
// 		nmk = "AB"
// 	}else if nam > 65 {
// 		nmk = "B"
// 	}else if nam > 57.7 {
// 		nmk = "BC"
// 	}else if nam > 50 {
// 		nmk = "C"
// 	}else if nam > 40 {
// 		nmk = "D"
// 	}else if nam <= 40 {
// 		nmk = "E"
// 	}
// 	fmt.Println("Nilai Mata Kuliah : ", nmk)
// }

// Jawaban A : Program tidak mengeluarkan output sesuai apa yang sudah di program alias error, Eksekusi tidak
// sesuai dengan soal
// Jawaban B : Kesalahan program terletak pada ekspresi percabangan seharusnya menggunakan variabel nmk
// sebagai variabel yang memiliki tipe data string. Kesalahan kedua terletak pada penggunaan percabangan itu sendiri
// awalnya menggunakan if namun jika ternyata percabangan lebih dari satu sebaiknya menggunakan else if.
// Jawaban C :

// Nomor 3
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var bil int
// 	fmt.Print("Bilangan : ")
// 	fmt.Scan(&bil)

// 	fmt.Printf("Faktor %d : ", bil)
// 	for i := 1; i <= bil; i++ {
// 		if bil % i == 0 {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// 	fmt.Println()

// 	isPrime := true
// 	for i := 2; i*i <= bil; i++{
// 		if bil %2 == 0{
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime{
// 		fmt.Println("Prima : ", isPrime)
// 	}else {
// 		fmt.Println("Prima : ", isPrime)
// 	}
// }
