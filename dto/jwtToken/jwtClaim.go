package jwttoken

type Claims struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	UserName string `json:"username"`
}
