package main

type User struct {
	Id       string   `json:"id,omitempty" form:"id"`
	Name     string   `json:"name,omitempty" form:"name"`
	Email    string   `json:"email,omitempty" form:"email"`
	Password string   `json:"-" form:"password"`
	Roles    []string `json:"roles"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}
