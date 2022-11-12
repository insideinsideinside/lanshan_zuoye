package dao

import "fmt"

var database = map[string]string{
	"yuanshen":  "123456",
	"wxgg":      "135246",
	"limingjie": "654321",
}

func Adduser(username, password string) {
	database[username] = password
}
func Selectuser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}
func Selectpassword(username string) string {
	return database[username]
}

func Changepassword(username, password string) {
	database[username] = password
	fmt.Println("更改成功")
}
