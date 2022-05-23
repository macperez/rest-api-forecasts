package controllers

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var jwtKey = []byte("my_secret_key")
