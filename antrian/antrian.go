package antrian

import (
	"fmt"
	"time"
)

func Antrian() {
	queue := make(chan string, 5)

	// inisiasi
	queue <- "andi"
	queue <- "Budi"
	queue <- "Cecep"
	queue <- "Doni"

	go serviceRoutine(queue)

	var input string
	for {
		fmt.Print("Masukan nama pendaftar : ")
		fmt.Scan(&input)

		if input == "exit" {
			close(queue)
			break
		}

		go serviceRegister(queue, input)
	}

	fmt.Println("Program end..")
}

func serviceRoutine(q chan string) {
	for nama := range q {
		fmt.Println("\nMelayani ", nama)
		time.Sleep(4 * time.Second)
	}
}

func serviceRegister(q chan string, nama string) {
	q <- nama
}
