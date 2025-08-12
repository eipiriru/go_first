package main

import "go_first/webapps"

// "go_first/kalkulator"

func main() {
	// for {
	// 	kalkulator.Kalkulator()
	// }

	address := "localhost:9000"
	webapps.Route(address)

	// webapps.SendRandomPoll("urlpolling", 2)
	// webapps.SendPollShouldBalanced("urlpolling")
	// webapps.SendPollSpecificUntilTarget("urlpolling", 0, 50)

	// antrian.Antrian()

	// fmt.Println(webapps.GetRandomKomentarByPrompt("Berikan satu quote dalam dari filsuf terkenal. Berikan jawaban dengan format kutipan diakhiri dengan pemberi kutipan"))

	// conf := c.Config()
	// loc, _ := time.LoadLocation(conf.TIMEZONE)
	// timeNow := time.Now().In(loc)

	// fmt.Printf("%v", timeNow)

	// newUser := cobasql.User{
	// 	Username:      "Testing",
	// 	Name:          "Testingm123",
	// 	Email:         "Testingm123@mail",
	// 	Register_date: &timeNow,
	// }
	// fmt.Println(newUser)
	// each := cobasql.InsertUser(newUser)
	// fmt.Println(each.Id, " - ", each.Username, " -", each.Name, " - ", each.Email, " - ", helpers.Ternary(each.Register_date != nil, each.Register_date, ""))

	// result := cobasql.GetAllUser()
	// for _, each := range result {
	// 	fmt.Println(each.Id, " - ", each.Username, " -", each.Name, " - ", each.Email, " - ", helpers.Ternary(each.Register_date != nil, each.Register_date, ""))
	// }

	// fmt.Printf("%v", time.Now())
}
