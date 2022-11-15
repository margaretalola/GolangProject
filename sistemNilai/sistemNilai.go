package main

import "fmt"

func main() {
	var grade int
	fmt.Print("Masukkan nilai anda: ")
	fmt.Scan(&grade)

	if grade >= 9 {
		fmt.Println("Anda mendapatkan nilai A.")
	} else if grade >= 6 {
		fmt.Println("Anda mendapatkan nilai B.")
	} else if grade >= 5 {
		fmt.Println("Anda mendapatkan nilai C.")
	} else {
		fmt.Println("Anda mendapatkan nilai E.")
	}
}

/*switch {
case grade >= 9:
	fmt.Println("Anda mendapatkan nilai A.")
case grade >= 6:
	fmt.Println("Anda mendapatkan nilai B.")
case grade >= 5:
	fmt.Println("Anda mendapatkan nilai C.")
default:
	fmt.Println("Anda mendapatkan nilai E.")
}*/
