package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Name string `json:"name"`
}

type EcClaims struct {
	jwt.StandardClaims
	User User `json:"user"`
}

const TokenExpireDuration = time.Hour * 2

var Secret = []byte("夏天夏天悄悄过去")

func GetExTime() int64 {
	return time.Now().Add(TokenExpireDuration).Unix()
}

func GetClaims(user User) EcClaims {
	c := EcClaims{
		User: user,
	}

	c.ExpiresAt = GetExTime()
	c.Issuer = "user"
	return c
}

// GenToken 生成JWT
func GenToken(c EcClaims) (string, error) {

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

func ParseToken(tokenString string) (*EcClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &EcClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*EcClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func RefreshToken(token string) (string, error) {
	var (
		newToken string
		err      error
	)
	c, err := ParseToken(token)
	c.ExpiresAt = GetExTime()
	newToken, err = GenToken(*c)
	return newToken, err

}

func main() {
	u := User{
		"123",
	}
	c := GetClaims(u)
	tk, _ := GenToken(c)
	fmt.Println(tk)
	m, e := ParseToken(tk)
	fmt.Println(e)
	fmt.Println(m.User.Name)
	fmt.Println(m.ExpiresAt)
	time.Sleep(time.Second * 2)
	tk, e = RefreshToken(tk)
	m, e = ParseToken(tk)
	fmt.Println(e)
	fmt.Println(m.ExpiresAt)
}
