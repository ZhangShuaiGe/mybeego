package dto

type UserInfo struct {
	Username string `json:"username" valid:"Required"`
	Password string `json:"password" valid:"Required"`
}
