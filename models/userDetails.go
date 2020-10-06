package models

var RegisteredUsers []UserDetials

type AuthResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

type UserDetials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
}
