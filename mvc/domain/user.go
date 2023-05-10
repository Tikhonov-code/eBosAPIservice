package domain

type User struct {
	Id        int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Role string `json:"role"`
	Status	string `json:"status"`
}

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}