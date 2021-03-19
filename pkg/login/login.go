package login

import (
	"github.com/dgrijalva/jwt-go"
)

type login struct {
	token  *jwt.Token
	Claims *Claims
}

type Claims struct {
	Sub string `json:"sub"`
	jwt.StandardClaims
}

func NewLogin(sub string) *login {
	var l login
	//	expirationTime := time.Now().Add(25 * time.Minute)

	l.Claims = &Claims{
		Sub:            sub,
		StandardClaims: jwt.StandardClaims{
			//			ExpiresAt: expirationTime.Unix(),
		},
	}

	l.token = jwt.NewWithClaims(jwt.SigningMethodHS256, l.Claims)
	return &l
}

func (l *login) GetToken(secret string) (string, error) {
	return l.token.SignedString([]byte(secret))
}
