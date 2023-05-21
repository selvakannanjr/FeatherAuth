package credmanage

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateSignedToken(username string,admin bool,expiry time.Time,secretkey string)(string,error){

	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretkey))
	return t,err
}

func HashGen(password string)string{

	hash,err := bcrypt.GenerateFromPassword([]byte(password),4)
	if err!=nil{
		return ""
	}

	return string(hash)
}

func ValidateHash(hash,givenpass string)bool{

	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(givenpass))
	return err == nil

}