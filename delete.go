package arangodb

import (
	"fmt"
	"github.com/apex/log"
	"github.com/thedanielforum/arangodb/errc"
)

func (c *Connection) Delete(collectionName string ,docHandle string) error {

	endPoint := fmt.Sprintf("_db/%s/_api/document/%s/%s", c.db, collectionName,docHandle)

	_,err := c.del(endPoint, nil)
	if err != nil {
		//log.WithError(err).Info(arangodb.ErrorCodeInvalidEdgeAttribute.Error().Error())
		return errc.ErrorCodeNoResult.Error()
	}
	if c.config.DebugMode {
		log.Infof("Deleted document %s in: %s", docHandle,collectionName)
	}
	return nil
}