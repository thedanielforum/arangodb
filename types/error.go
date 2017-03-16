package types

import (
	"errors"
	"encoding/json"
	"github.com/apex/log"
)

type DbError struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
	ErrorNum     int    `json:"errorNum"`
}

func NewDbError(body []byte) *DbError {
	db := new(DbError)
	err := json.Unmarshal(body, db)
	if err != nil {
		// I don't call fatal as it would shutdown the parent app using the driver.
		log.WithError(err).Error("error while unmarshaling error response from server")
		return nil
	}
	return db
}

func (e *DbError) ToError() error {
	return errors.New(e.ErrorMessage)
}


