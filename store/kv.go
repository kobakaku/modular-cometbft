package store

import (
	"path"
	"strings"

	ds "github.com/ipfs/go-datastore"

	badger4 "github.com/ipfs/go-ds-badger4"
)

func NewKVStore(dbName string) (ds.TxnDatastore, error) {
	return badger4.NewDatastore(dbName, nil)
}

// GenerateKey creates a key from a slice of string fields, joining them slashes.
func GenerateKey(fields []string) string {
	key := "/" + strings.Join(fields, "/")
	return path.Clean(key)
}
