package define

import "github.com/dgrijalva/jwt-go"

const Port = "40010"
const Username = "root"
const Password = "123456"
const Dbname = "terminal"
const JwtSecret = "terminal-key"

const BucketPath = "https://terminal-1304032890.cos.ap-nanjing.myqcloud.com"
const SecretID = "AKIDwjL5btxnPnwsb5MPUL0TkodEfsrb6Zpq"
const SecretKey = "sWw73lFVXO6apppcEeuynOlvsc8SYRTP"

type UserClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
