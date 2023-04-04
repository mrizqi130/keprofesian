package main

import "fmt"

func main() {
	var jumlah,ganjil,genap int
	fmt.Println("Program menentukan angka ganjil dan genap dari 1 hingga n")
	fmt.Print("Masukkan angka n : ")
	fmt.Scanln(&jumlah)
	for i := 1; i <= jumlah; i++ {
		if i%2==0 {
			fmt.Println(i," adalah angka genap")
			genap++
		} else {
			fmt.Println(i," adalah angka ganjil")
			ganjil++
		}
	}
	fmt.Println()
	fmt.Println("Angka ganjil berjumlah :", ganjil)
	fmt.Println("Angka genap berjumlah :", genap)
}