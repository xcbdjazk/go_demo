package user

type User struct {
	Name     string
	PassWord string
}

func (user *User) SetUserPassWord(newPassWord string) {
	user.PassWord = newPassWord
}

type UserSlice []User

var Users UserSlice

func init() {
	u1 := User{"jack", "123123"}
	u2 := User{"zhangshan", "123456"}
	u3 := User{"mary", "123123"}
	Users = append(Users, u1)
	Users = append(Users, u2)
	Users = append(Users, u3)
}

func GetUser(name, pwd string) (User, bool) {
	for _, v := range Users {
		if v.Name == name && pwd == v.PassWord {
			return v, true
		}
	}
	return User{}, false
}

func GetUserByName(name string) (user User, err bool) {
	for _, user = range Users {
		if user.Name == name {
			return user, true
		}
	}
	return
}

func AddUser(name, pwd string) (User, bool) {
	if _, ok := GetUserByName(name); ok {
		return User{}, false
	}
	u := User{name, pwd}
	Users = append(Users, u)
	return u, true
}
