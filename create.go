package arangodb

import (
	"encoding/json"
	"github.com/apex/log"
	"fmt"
	"github.com/thedanielforum/arangodb/errc"
)

func (c *Connection) Create(col string, doc interface{}) (*Document, error) {
	responseDoc := &Document{}

	// Check if we need to create a new collection
	if c.config.AutoCreateColOnInsert {
		c.cacheValidation(col, doc)
	}

	// Collection Confirm Exist Now , Proceed to perform save document/edge
	url := fmt.Sprintf("_db/%s/_api/document/%s", c.db, col)
	encoded,err := json.Marshal(doc)
	if err != nil {
		return responseDoc, err
	}

	resp, err := c.post(url, encoded)
	if err != nil {
		if err.Error() == errc.ErrorCodeInvalidEdgeAttribute.String() {
			log.WithError(err).Info(errc.ErrorCodeInvalidEdgeAttribute.Msg())
			return responseDoc, err
		}
		log.WithError(err).Info(err.Error())
		return responseDoc, err
	}

	err = json.Unmarshal(resp, &responseDoc)
	if err != nil {
		return responseDoc, err
	}
	if c.config.DebugMode {
		log.Infof("Created document in: %s", col)
	}
	return responseDoc, nil
}

// cacheValidation checks internal cache if such collection exist before attempting to create new collection
func (c *Connection) cacheValidation(collectionName string, doc interface{}) error {
	// true means that collection exist
	if c.colCache[c.db][collectionName] {
		return nil
	}
	// if collection don't exist, create one
	checkEdge := new(Document)
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
