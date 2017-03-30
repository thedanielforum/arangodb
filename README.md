# arangodb driver

How to install
```bash
go get github.com/thedanielforum/arangodb
```

Set up a connection to arango
```go
package main

import "github.com/thedanielforum/arangodb"

var conn *arangodb.Connection

func main()  {
	var err error
	conn, err = arangodb.NewConnection(
		"127.0.0.1:8529",
		"root",
		"awstest123",
		&arangodb.Config{
			DebugMode: false,
		},
	)
	conn.SetDB("test")
	if err != nil {
		panic(err)
	}
}
```

Get one document
```go
var vid *Videos
err = connection.NewQuery(`FOR x IN %s RETURN x`, "videos").One(&vid)
if err != nil {
	panic(err)
}
```

Get all documents from query
```go
var vids []Video
err = connection.NewQuery(`FOR x IN %s RETURN x`, "videos").All(&vids)
if err != nil {
	panic(err)
}
```
