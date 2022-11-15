//2.	Buat program percabangan dimana saat melakukan input A
//dia akan melakukan perulangan sebanyak 10x berisi Nama dan
//bila di input B mengulang NPM kalian (10 Point)

package main

import "fmt"

func main() {
	var choice string
	fmt.Print("Masukkan Input (a/b): ")
	fmt.Scan(&choice)
	var i int
	if choice == "a" {
		for i = 1; i <= 10; i++ {
			fmt.Println("Nama ke-", i)
		}
	} else if choice == "b" {
		for i = 1; i <= 10; i++ {
			fmt.Println("NPM ke-", i)
		}
	} else {
		fmt.Println("Maaf masukkan anda salah.")
	}
}
