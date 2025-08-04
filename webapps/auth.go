package webapps

const USER = "admin"
const PASS = "pass123"

func auth(username string, password string) bool {
	return (username == USER) && (password == PASS)
}
