package webapps

import "go_first/cobasql"

func GetAllUsers() []cobasql.User {
	return cobasql.GetAllUser()
}
