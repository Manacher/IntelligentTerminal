package request

type UserRegisterReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Name     string `json:"name"`
}
