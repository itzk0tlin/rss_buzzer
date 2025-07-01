package pkg

import (
	"bytes"
	"os"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
)

type DefaultRepository struct {
	path string
}

func (repo DefaultRepository) OpenDatabase() (*badger.DB, error) {
	_, err := os.Stat(repo.path)
	if os.IsNotExist(err) {
		wd, _ := os.Getwd()
		path := string(wd) + string(os.PathSeparator) + repo.path
		os.Mkdir(path, 0755)
	}
	opts := badger.DefaultOptions(repo.path)
	opts.Compression = options.ZSTD
	opts.ZSTDCompressionLevel = 2
	db, err := badger.Open(opts)
	return db, err
}

func (repo DefaultRepository) GetValue(key []byte) ([]byte, error) {
	db, err := repo.OpenDatabase()
	var resultValue []byte
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {
			resultValue = val
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resultValue, nil
}

func (repo DefaultRepository) GetKey(value []byte) ([]byte, error) {
	db, err := repo.OpenDatabase()
	var resultKey []byte
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		opts.Prefix = nil
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			item.Value(func(val []byte) error {
				if bytes.Equal(val, value) {
					resultKey = key
				}
				return nil
			})
		}
		return nil
	})
	return resultKey, nil
}

func (repo DefaultRepository) DeleteValue(key []byte) error {
	db, err := repo.OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (repo DefaultRepository) InsertValue(key []byte, value []byte) error {
	db, err := repo.OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
