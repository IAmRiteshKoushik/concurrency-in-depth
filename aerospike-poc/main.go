package main

import (
	"fmt"

	aero "github.com/aerospike/aerospike-client-go/v8"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := aero.NewClient("localhost", 3000)
	panicOnError(err)

	key, err := aero.NewKey("test", "aerospike", "key")
	panicOnError(err)

	bins := aero.BinMap{
		"bin1": 42,
		"bin2": "An elephant is a mouse with an ops sys",
		"bin3": []any{"Go", 2009},
	}

	err = client.Put(nil, key, bins)
	panicOnError(err)

	rec, err := client.Get(nil, key)
	panicOnError(err)

	existed, err := client.Delete(nil, key)
	panicOnError(err)
	fmt.Printf("Record existed before delete? %v\n", existed)
}
