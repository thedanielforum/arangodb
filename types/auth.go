package types

type Auth struct {
	Jwt                string `json:"jwt"`
	MustChangePassword bool   `json:"must_change_password"`
}
