package types

type DbInfo struct {
	Result struct {
		Name     string `json:"name"`
		ID       string `json:"id"`
		Path     string `json:"path"`
		IsSystem bool   `json:"isSystem"`
	}          `json:"result"`
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

//Get properties of collection
type CollectionInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsSystem bool   `json:"isSystem"`
	Type     int    `json:"type"`
	Code     int    `json:"code"`
}

//Get properties of all collections
type ColInfo struct {
	Result []CollectionInfo `json:"result"`
	Error  bool             `json:"error"`
	Code   int              `json:"code"`
}