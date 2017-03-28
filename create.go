package arangodb

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
)

//type 2 = document collection
//type 3 = edge collection
//name = collection of name
type CollectionProp struct {
	JournalSize uint                   `json:"journalSize,omitempty"`
	Name        string                 `json:"name"`
	Type        uint                   `json:"type"`
	WaitForSync bool                   `json:"waitForSync,omitempty"`
	isVolatile  bool                   `json:"isVolatile,omitempty"`
	Shards      int                    `json:"numberOfShards,omitempty"`
	ShardKeys   []string `              json:"shardKeys,omitempty"`
	Keys        map[string]interface{} `json:"keyOptions,omitempty"`
}

func (c *Connection) Create(typeCollection uint,collections ...string) error {
	//any other number in typeCollection apart from 2&3 will lead to 2 by default
	for _, collection := range collections {
		preFab, err := json.Marshal(&CollectionProp{
			Name: collection,
			Type: typeCollection,
		})
		if err != nil {
			return err
		}

		endPoint := fmt.Sprintf("/_db/%s/_api/collection",c.db)
		_, err = c.post(endPoint,preFab)
		if err != nil {
			log.WithError(err).Info("Collection Already Exist")
			return err
		}

	}
	return nil
}

func (c *Connection) Save(collectionName string, doc interface{}) error {
	endPoint := fmt.Sprintf("/_db/%s/_api/document/%s",c.db,collectionName)
	preFab,err := json.Marshal(doc)
	if err != nil {
		return err
	}

	_, err = c.post(endPoint,preFab)
	if err != nil {
		return nil
	}
	return nil
}