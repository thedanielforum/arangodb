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

