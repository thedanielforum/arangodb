package arangodb

import (
	"encoding/json"
	"github.com/apex/log"
	"fmt"
	"github.com/thedanielforum/arangodb/types"
	"github.com/thedanielforum/arangodb/errc"
)

type EdgeProp struct {
	From      string    `json:"_from"`
	To        string    `json:"_to"`
}

func (c *Connection) Create(col string, doc interface{}) error {
	c.cacheValidation(col, doc)

	// Collection Confirm Exist Now , Proceed to perform save document/edge
	url := fmt.Sprintf("/_db/%s/_api/document/%s", c.db, col)
	encoded,err := json.Marshal(doc)
	if err != nil {
		return err
	}
	_, err = c.post(url, encoded)
	if err != nil {
		if err.Error() == errc.ErrorCodeInvalidEdgeAttribute.String() {
			log.WithError(err).Info(errc.ErrorCodeInvalidEdgeAttribute.Msg())
			return err
		}
		log.WithError(err).Info(err.Error())
		return err
	}
	return nil
}

// cacheValidation checks internal cache if such collection exist before attempting to create new collection
func (c *Connection) cacheValidation(collectionName string, doc interface{}) error {
	// true means that collection exist
	if c.colCache[c.db][collectionName] {
		return nil
	}
	// if collection don't exist, create one
	checkEdge := new(EdgeProp)
	encodedDoc, _ := json.Marshal(doc)
	json.Unmarshal(encodedDoc, checkEdge)

	if checkEdge.To != "" && checkEdge.From != "" {
		c.NewEdge(collectionName)
		cacheAdd(c.colCache, c.db, collectionName)
		return nil
	}
	c.NewCollection(collectionName)
	cacheAdd(c.colCache, c.db, collectionName)
	return nil
}

func (c *Connection) GetAllColProp() {
	urlStack := fmt.Sprintf("/_db/%s/_api/collection",c.db)

	//err means that database do not exist
	allProp, err := c.get(urlStack)
	if err != nil {
		log.WithError(err).Info(errc.ErrorCodeNoDatabaseSelected.Msg())
	}
	colsInfo := new(types.ColInfo)
	json.Unmarshal(allProp, colsInfo)

	for _, colInfo := range colsInfo.Result{
		if colInfo.IsSystem == false {
			cacheAdd(c.colCache, c.db, colInfo.Name)
		}
	}
	return
}