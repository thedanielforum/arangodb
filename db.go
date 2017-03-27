package arangodb

import (
	"github.com/apex/log"
	"github.com/thedanielforum/arangodb/types"
	"fmt"
	"encoding/json"
)

func (c *Connection) SetDB(db string) *Connection {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.db = db
	if c.config.DebugMode {
		log.Infof("selected db: %s", c.db)
	}
	return c
}

func (c *Connection) DBInfo() (*types.DbInfo, error) {
	var err error
	info := new(types.DbInfo)

	if c.db == "" {
		return info, ErrorCodeNoDatabaseSelected.Error()
	}

	body, err := c.get(fmt.Sprintf("_db/%s/_api/database/current", c.db))
	if err != nil {
		return info, err
	}

	if err = json.Unmarshal(body, info); err != nil {
		return info, err
	}

	return info, nil
}

func (c *Connection) ListDBs() (*types.Dbs, error) {
	var err error
	dbs := new(types.Dbs)

	if c.db == "" {
		return dbs, ErrorCodeNoDatabaseSelected.Error()
	}

	body, err := c.get("_api/database")
	if err != nil {
		return dbs, err
	}

	if err = json.Unmarshal(body, dbs); err != nil {
		return dbs, err
	}

	return dbs, nil
}

