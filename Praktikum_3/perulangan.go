// ARVEL MAHSA 
//SOAL LATIHAN 1
// package main

// import ("fmt")

// func main (){
// 	//Definisikan warna yang benar 
// 	urutanBenar := []string{"merah","kuning","hijau","ungu"}
// 	hasil := true

// 	//Lakukan 5 percobaan
// 	for i := 1; i <= 5; i++ {
// 		var warna1, warna2, warna3, warna4 string
// 		fmt.Printf("Percobaan %d\n", i)
// 		fmt.Print("Masukan warna pertama : ")
// 		fmt.Scanln(&warna1)
// 		fmt.Print("Masukan warna kedua : ")
// 		fmt.Scanln(&warna2)
// 		fmt.Print("Masukan warna ketiga : ")
// 		fmt.Scanln(&warna3)
// 		fmt.Print("Masukan warna keempat : ")
// 		fmt.Scanln(&warna4)

// 		//Periksa apakah urutan warna sesuai
// 		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] ||
// 		warna3 != urutanBenar[2] || warna4 != urutanBenar[3]{
// 			hasil = false
// 		}
// 	}
// 	// Tampilkan hasil
// 	fmt.Println("Berhasil : ", hasil)
// }

// SOAL LATIHAN 2
// package main

// import ("fmt" 
// "bufio" 
// "os"
// "strings"
// )
// func main () {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var pita string
// 	var bungaCount int

// 	for{
// 		fmt.Printf("Bunga %d: ", bungaCount + 1)
// 		scanner.Scan()
// 		input := scanner.Text()

// 		if strings.ToLower(input) == "selesai" {
// 			break
// 		}
// 		if pita == "" {
// 			pita = input
// 		}else {
// 			pita += " " + input
// 		}
// 		bungaCount++
// 	}

// 	fmt.Println("Pita : ", pita)
// 	fmt.Println("Bunga : ", bungaCount)

// }


