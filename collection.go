package arangodb

import (
	"fmt"
	"encoding/json"
	"github.com/apex/log"
	"github.com/thedanielforum/arangodb/errc"
	"github.com/thedanielforum/arangodb/types"
)

//type 2 = document collection
//type 3 = edge collection
const (
	TypeDoc = iota + 2
	TypeEdge
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

func (c *Connection) NewCollection(cols ...string) {
	for _, col := range cols {
		preFab, err := json.Marshal(&CollectionProp{
			Name: col,
			Type: TypeDoc,
		})
		if err != nil {
			log.WithError(err).Error("encoding error")
		}
		endPoint := fmt.Sprintf("_db/%s/_api/collection", c.db)
		_, err = c.post(endPoint, preFab)
		if err != nil {
			log.Infof("Document collection already exist: %s", col)
		} else {
			if c.config.DebugMode {
				log.Infof("Created document collection: %s", col)
			}
		}
	}
}

func (c *Connection) NewEdge(edges ...string) {
	for _, edge := range edges {
		preFab, err := json.Marshal(&CollectionProp{
			Name: edge,
			Type: TypeEdge,
		})
		if err != nil {
			log.WithError(err).Error("encoding error")
		}

		endPoint := fmt.Sprintf("_db/%s/_api/collection", c.db)
		_, err = c.post(endPoint, preFab)
		if err != nil {
			log.Infof("Edge collection already exist: %s", edge)
		} else {
			if c.config.DebugMode {
				log.Infof("Created edge collection: %s", edge)
			}
		}
	}
}

func (c *Connection) GetAllCollections() error {
	url := fmt.Sprintf("_db/%s/_api/collection", c.db)

	// err means that database do not exist
	resp, err := c.get(url)
	if err != nil {
		log.WithError(err).Info(errc.ErrorCodeNoDatabaseSelected.Msg())
		return errc.ErrorCodeNoDatabaseSelected.Error()
	}
	cols := new(types.ColInfo)
	err = json.Unmarshal(resp, cols)
	if err != nil {
		return err
	}

	for _, col := range cols.Result {
		if col.IsSystem == false {
			cacheAdd(c.colCache, c.db, col.Name)
		}
	}
	return nil
}
