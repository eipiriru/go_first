package cobasql

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	Id            int64
	Username      string
	Name          string
	Email         string
	Register_date *time.Time
}

func GetAllUser() []User {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	sql := "SELECT Id, IFNULL(Username, ''), IFNULL(Name, ''), IFNULL(Email, ''), Register_date FROM users LIMIT 5"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	result := []User{}
	for rows.Next() {
		each := User{}
		err := rows.Scan(&each.Id, &each.Username, &each.Name, &each.Email, &each.Register_date)
		if err != nil {
			fmt.Println(err.Error())
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return []User{}
	}
	return result
}

func GetUserById(id int64) User {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	result := User{}
	sql := "SELECT Id, IFNULL(Username, ''), IFNULL(Name, ''), IFNULL(Email, ''), Register_date FROM users WHERE id=?"
	err = db.QueryRow(sql, id).Scan(&result.Id, &result.Username, &result.Name, &result.Email, &result.Register_date)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func InsertUser(u User) User {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	sql := "INSERT INTO users (username, name, email, register_date) VALUES (?,?,?,?)"
	res, err := db.Exec(sql, u.Username, u.Name, u.Email, u.Register_date)
	if err != nil {
		fmt.Println(err.Error())
		return User{}
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}

	return GetUserById(lastId)
}

func UpdateUser(u User) User {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	sql := "UPDATE users SET username=?, name=?, email=? WHERE id=?"
	_, err = db.Exec(sql, u.Username, u.Name, u.Email, u.Id)
	if err != nil {
		fmt.Println(err.Error())
		return User{}
	}

	return GetUserById(u.Id)
}

func DeleteUserById(id int64) bool {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	sql := "DELETE FROM users WHERE id=?"
	_, err = db.Exec(sql, id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
