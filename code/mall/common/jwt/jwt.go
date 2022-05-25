package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(secretKey string, issueAt, sec, uid int64) (string, error) {
	claim := make(jwt.MapClaims)
	//Claim field
	claim["exp"] = issueAt + sec
	claim["iat"] = issueAt
	claim["uid"] = uid

	//create jwt
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim                         //assign claim to token
	return token.SignedString([]byte(secretKey)) //sign token with key
}
