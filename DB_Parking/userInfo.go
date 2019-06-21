package DB_Parking

type User struct {
	UserName string
	Crippled     string
}

var UserMap = make(map[string]User)


func AddNewUser(user User) {
	UserMap[user.UserName] = user
}

func UserExists(userName string) bool {
	_, validUser := UserMap[userName]
	if validUser {
		return true
	}
	return false
}
