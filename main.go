package main

import "go_first/webapps"

// "go_first/kalkulator"

func main() {
	// for {
	// 	kalkulator.Kalkulator()
	// }

	address := "localhost:9000"
	webapps.Route(address)
}
