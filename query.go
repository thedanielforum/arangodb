package arangodb

import (
	"fmt"
	"strings"
	"github.com/thedanielforum/arangodb/types"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/thedanielforum/arangodb/errc"
)

type Query struct {
	aql        string
	bindParams map[string]interface{}
	cache      bool
	batchSize  int

	conn       *Connection
}

func (c *Connection) NewQuery(aql string, params ...interface{}) *Query {
	// Process to remove eventual tabs/spaces used when indenting the query
	aql = cleanAQLQuery(aql)
	aql = fmt.Sprintf(aql, params...)
	// Replace by single quotes so there is no conflict when serialised in JSON
	aql = strings.Replace(aql, `"`, "'", -1)

	// TODO Batch size default
	return &Query{
		aql:       aql,
		batchSize: 100,
		conn:      c,
	}
}

// Cache enables/disables the caching of the query.
// Unavailable prior to ArangoDB 2.7
func (q *Query) Cache(enable bool) *Query {
	q.cache = enable
	return q
}

// BatchSize sets the batch size of the query
func (q *Query) BatchSize(size int) *Query {
	q.batchSize = size
	return q
}

func (q *Query) One(result interface{}) (err error) {
	aql, err := json.Marshal(&types.Query{
		Aql:       q.aql,
		Count:     true,
		BatchSize: 1,
	})
	if err != nil {
		return err
	}

	resp, err := q.conn.post(fmt.Sprintf("/_db/%s/_api/cursor", q.conn.db), aql)
	if err != nil {
		return err
	}

	reply := new(types.Result)
	err = json.Unmarshal(resp, reply)
	if err != nil {
		return err
	}

	// Check for DB error
	if reply.Error {
		return errors.New("Query error")
	}

	// Check for no results
	if (len(reply.Result) <= 0) {
		return errc.ErrorCodeNoResult.Error()
	}

	// Parse reply in to result interface{}
	err = json.Unmarshal(reply.Result[0], result)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) All(result interface{}) (err error) {
	aql, err := json.Marshal(&types.Query{
		Aql:       q.aql,
		Count:     true,
		BatchSize: q.batchSize,
	})
	if err != nil {
		return err
	}

	resp, err := q.conn.post(fmt.Sprintf("/_db/%s/_api/cursor", q.conn.db), aql)
	if err != nil {
		return err
	}

	reply := new(types.Results)
	err = json.Unmarshal(resp, reply)
	if err != nil {
		return err
	}

	// Check for DB error
	if reply.Error {
		return errors.New("Query error")
	}

	// Check for no results
	if (len(reply.Result) <= 0) {
		return errc.ErrorCodeNoResult.Error()
	}

	// Parse reply in to result interface{}
	err = json.Unmarshal(reply.Result, &result)
	if err != nil {
		return err
	}

	return nil
}

// TODO Optimize
func cleanAQLQuery(aql string) string {
	aql = strings.Replace(aql, "\n", " ", -1)
	aql = strings.Replace(aql, "\t", "", -1)

	split := strings.Split(aql, " ")
	split2 := []string{}

	for _, s := range split {
		if len(s) == 0 {
			continue
		}
		split2 = append(split2, s)
	}

	return strings.Join(split2, " ")
}
