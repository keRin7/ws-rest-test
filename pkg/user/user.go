package user

import (
	"math/rand"
	"strconv"
	"time"

	guuid "github.com/google/uuid"
)

type User struct {
	tel   string
	uuid  string
	token string
}

func NewUser() *User {
	var c User

	c.uuid = guuid.New().String()

	min := 1000000
	max := 9999999
	rand.Seed(time.Now().UnixNano())
	c.tel = "+7929" + strconv.Itoa(rand.Intn(max-min)+min)

	return &c
}

func (u *User) GetTel() string {
	return u.tel
}

func (u *User) GetUUID() string {
	return u.uuid
}

func (u *User) SetToken(token string) {
	u.token = token
}

func (u *User) GetToken() string {
	return u.token
}
