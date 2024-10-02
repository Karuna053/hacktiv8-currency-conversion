package main

import (
	"fmt"
)

func main() {
	// Deklarasi nilai tukar
	var rupiah float64
	fmt.Println("Input how much rupiah do you have: ")
	fmt.Scanln(&rupiah)

	// currency disini
	var usdToIdr float64 = 15000
	var eurToIdr float64 = 16000
	var jpyToIdr float64 = 140

	fmt.Println("IDR to USD: ", rupiah/usdToIdr)
	fmt.Println("IDR to EUR: ", rupiah/eurToIdr)
	fmt.Println("IDR to JPY: ", rupiah/jpyToIdr)

	fmt.Println("Press Enter to exit program!")
	fmt.Scanln() // Wait for input
}
