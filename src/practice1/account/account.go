package account

import (
	"fmt"
	"practice1/user"
)

type Account struct {
	Money float64
}
type Account2User struct {
	Account *Account
	User    *user.User
}

type Account2UserSlice []Account2User

var Aus Account2UserSlice

func GetAU(user *user.User) *Account {
	for _, v := range Aus {
		if v.User == user {
			return v.Account
		}
	}
	au := Account2User{
		Account: &Account{},
		User:    user,
	}
	Aus = append(Aus, au)
	return au.Account
}
func SaveMoney(u *user.User, money float64) {
	au := GetAU(u)
	au.Money += money
}

func GetMoney(u *user.User, money float64) {
	au := GetAU(u)
	if au.Money > money {
		au.Money -= money
		fmt.Printf("你取出了，%v \n", money)
	} else {
		fmt.Printf("钱不够\n")
	}

}
