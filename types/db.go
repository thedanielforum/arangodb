package types

type DbInfo struct {
	Result struct {
		       Name     string `json:"name"`
		       ID       string `json:"id"`
		       Path     string `json:"path"`
		       IsSystem bool   `json:"isSystem"`
	       }   `json:"result"`
	Error bool `json:"error"`
	Code  int  `json:"code"`
}

type Dbs struct {
	Result []string `json:"result"`
	Error  bool     `json:"error"`
	Code   int      `json:"code"`
}

type Db struct {
	Name  string `json:"name"`
	Users []User `json:"users"`
}
