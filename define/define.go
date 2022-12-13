package define

import "github.com/dgrijalva/jwt-go"

const Port = "40010"

const Username = "root"
const Password = "123456"

// const Username = "debian-sys-maint"
// const Password = "xAftDU4Lgms6LLFJ"
const Dbname = "terminal"
const JwtSecret = "terminal-key"

const BucketPath = "https://terminal-1304032890.cos.ap-nanjing.myqcloud.com"
const SecretID = "AKIDwjL5btxnPnwsb5MPUL0TkodEfsrb6Zpq"
const SecretKey = "sWw73lFVXO6apppcEeuynOlvsc8SYRTP"
const DefaultAvatar = "https://terminal-1304032890.cos.ap-nanjing.myqcloud.com/terminal/default.png"

const FollowerListPageSize = 15

const MomentPageSize = 10
const MomentCommentPageSize = 10
const MomentSubCommentPageSize = 15

type UserClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
