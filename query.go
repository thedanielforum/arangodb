package arangodb

import "fmt"

type Query struct {
	query string
}

func (c *Connection) NewQuery(query string, params ...interface{}) *Query {
	q := new(Query)
	q.query = fmt.Sprintf(query, params)
	println(q.query)
	return q
}
