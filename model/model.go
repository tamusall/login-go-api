package model

type User struct {
	Email    string `json:email`
	Password string `json:password`
	Token    string `json:token`
}

type LoginResult struct {
	Error  string `json:error`
	Result string `json:result`
}
