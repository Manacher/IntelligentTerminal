package define

import "github.com/dgrijalva/jwt-go"

const Port = "40010"
const Username = "root"
const Password = "123456"
const Dbname = "terminal"
const JwtSecret = "terminal-key"

type UserClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
