package types

type User struct {
	Username string `json:"username"`
	Password string `json:"passwd"`
	Active   bool   `json:"active"`
}

