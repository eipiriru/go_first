package kalkulator

import (
	"fmt"
	"os"
)

func Kalkulator() {
	for {
		menuFunc()
	}
}

func menuFunc() {
	fmt.Println("== MENU ==")
	fmt.Println("1. Kalkulator Sederhana")
	fmt.Println("2. Konversi Suhu")
	fmt.Println("0. Exit Program")

	var choice int
	fmt.Print("Pilih menu : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		kalkulatorFunc()
	case 2:
		konversiSuhuFunc()
	case 0:
		fmt.Print("Exit Program")
		os.Exit(0)
	default:
		return
	}
}

func kalkulatorFunc() {
	fmt.Println("== Kalkulator Sederhana ==")
	fmt.Println("1. Perkalian")
	fmt.Println("2. Pembagian")

	var choice int
	fmt.Print("Pilih menu : ")
	fmt.Scan(&choice)

	var num1, num2 float64
	fmt.Print("Masukan angka 1 : ")
	fmt.Scan(&num1)
	fmt.Print("Masukan angka 2 : ")
	fmt.Scan(&num2)

	switch choice {
	case 1:
		fmt.Printf("Hasil %.2f * %.2f = %.2f \n", num1, num2, num1*num2)
	case 2:
		if num2 == 0 {
			fmt.Println("Tidak bisa melakukan pembagian dengan 0")
			kalkulatorFunc()
		} else {
			fmt.Printf("Hasil %.2f / %.2f = %.2f \n", num1, num2, num1/num2)
		}
	default:
		return
	}
}

func konversiSuhuFunc() {
	var listSuhu = []string{"Celcius", "Fahrenheit", "Kelvin", "Reamur"}

	showListSuhu := func(ls []string) {
		for i, s := range ls {
			fmt.Printf("%d. %s\n", i+1, s)
		}
	}
	fmt.Println("== Konversi Suhu ==")

	var suhu float64
	fmt.Print("Besar suhu : ")
	fmt.Scan(&suhu)

	showListSuhu(listSuhu)
	var dariSuhu int
	fmt.Print("Dari satuan suhu : ")
	fmt.Scan(&dariSuhu)

	showListSuhu(listSuhu)
	var keSuhu int
	fmt.Print("Ke satuan suhu : ")
	fmt.Scan(&keSuhu)

	var hasilKonversi float64

	convertSuhu := func(dari int, ke int, suhu float64) float64 {
		var suhuCelcius float64
		var result float64
		switch dari {
		case 1:
			suhuCelcius = suhu
		case 2:
			suhuCelcius = (5.0 / 9.0) * (suhu - 32)
		case 3:
			suhuCelcius = suhu - 273.15
		case 4:
			suhuCelcius = suhu * (5.0 / 4.0)
		}

		switch ke {
		case 1:
			result = suhuCelcius
		case 2:
			result = (suhuCelcius * (9.0 / 5.0)) + 32
		case 3:
			result = suhuCelcius + 273.15
		case 4:
			result = suhuCelcius * (4.0 / 5.0)
		}

		return result
	}

	hasilKonversi = convertSuhu(dariSuhu, keSuhu, suhu)
	fmt.Printf("Konversi -> %.2f %s = %.2f %s\n", suhu, listSuhu[dariSuhu-1], hasilKonversi, listSuhu[keSuhu-1])
}
