package pkg

import "github.com/dgraph-io/badger/v4"

type DBPair struct {
	Key   []byte
	Value []byte
}

type Repository interface {
	OpenDatabase() (*badger.DB, error)
	GetValue(key []byte) ([]byte, error)
	GetKey(value []byte) ([]byte, error)
	GetAllPairs() []DBPair
	DeleteValue(key []byte) error
	InsertValue(key []byte, value []byte) error
}
