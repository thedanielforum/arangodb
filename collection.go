package arangodb

import (
	"fmt"
	"encoding/json"
	"github.com/apex/log"
)

//type 2 = document collection
//type 3 = edge collection
const (
	TypeDoc = 2
	TypeEdge = 3
)

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

func (c *Connection) CreateColDoc(collections ...string) error {
	//any other number in typeCollection apart from 2&3 will lead to 2 by default
	for _, collection := range collections {
		preFab, err := json.Marshal(&CollectionProp{
			Name: collection,
			Type: TypeDoc,
		})
		if err != nil {
			return err
		}

		endPoint := fmt.Sprintf("/_db/%s/_api/collection",c.db)
		_, err = c.post(endPoint,preFab)
		if err != nil {
			log.WithError(err).Info(collection + " Document Collection Already Exist")
		}

	}
	return nil
}

func (c *Connection) CreateColEdge(collections ...string) error {
	//any other number in typeCollection apart from 2&3 will lead to 2 by default
	for _, collection := range collections {
		preFab, err := json.Marshal(&CollectionProp{
			Name: collection,
			Type: TypeEdge,
		})
		if err != nil {
			return err
		}

		endPoint := fmt.Sprintf("/_db/%s/_api/collection",c.db)
		_, err = c.post(endPoint,preFab)
		if err != nil {
			log.WithError(err).Info(collection + " Edge Collection Already Exist")
		}

	}
	return nil
}