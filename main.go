package main

import "go_first/webapps"

// "go_first/kalkulator"

func main() {
	// for {
	// 	kalkulator.Kalkulator()
	// }

	// address := "localhost:9000"
	// webapps.Route(address)

	// webapps.SendRandomPoll("urlpolling", 20)
	// webapps.SendPollShouldBalanced("urlpolling")
	webapps.SendPollSpecificUntilTarget("urlpolling", 0, 50)

	// antrian.Antrian()
}
