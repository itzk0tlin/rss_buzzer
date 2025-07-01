package pkg

import (
	"bytes"
	"log"
	"os"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
)

var RSSRepo = RSSRepository{Path: "data"}

type RSSRepository struct {
	Path string
}

func (repo RSSRepository) OpenDatabase() *badger.DB {
	_, err := os.Stat(repo.Path)
	if os.IsNotExist(err) {
		log.Println("Database not found creating.")
		wd, _ := os.Getwd()
		path := string(wd) + string(os.PathSeparator) + repo.Path
		os.Mkdir(path, 0755)
		log.Println("Database path created.")
	}
	opts := badger.DefaultOptions(repo.Path)
	opts.Compression = options.ZSTD
	opts.ZSTDCompressionLevel = 2
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Encountered critical error. Database failed to open: %s!\n", err)
	}
	log.Println("Database opened.")
	return db
}

func (repo RSSRepository) GetValue(key []byte) ([]byte, error) {
	key = []byte("feeds:" + string(key))
	db := repo.OpenDatabase()
	var resultValue []byte
	defer db.Close()
	log.Printf("Looking in database for value with key: %s.\n", key)
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			log.Printf("Database encountered error: %s!\n", err)
			return err
		}
		item.Value(func(val []byte) error {
			resultValue = val
			log.Printf("Get value %s from database.\n", val)
			return nil
		})
		return nil
	})
	if err != nil {
		log.Printf("Database encountered error: %s!\n", err)
		return nil, err
	}
	return resultValue, nil
}

func (repo RSSRepository) GetKey(value []byte) ([]byte, error) {
	db := repo.OpenDatabase()
	var resultKey []byte
	defer db.Close()
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		opts.Prefix = []byte("feeds:")
		it := txn.NewIterator(opts)
		defer it.Close()
		log.Printf("Looking for key in database with value %s.\n", value)
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			item.Value(func(val []byte) error {
				if bytes.Equal(val, value) {
					resultKey = key
					log.Printf("Get key %s from database.\n", key)
				}
				return nil
			})
		}
		return nil
	})
	return resultKey, nil
}

func (repo RSSRepository) GetAllPairs() []DBPair {
	db := repo.OpenDatabase()
	defer db.Close()
	var resultPairs []DBPair
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		opts.Prefix = []byte("feeds:")
		it := txn.NewIterator(opts)
		defer it.Close()
		log.Printf("Parsing all pairs...\n")
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			item.Value(func(val []byte) error {
				pair := DBPair{Key: key, Value: val}
				resultPairs = append(resultPairs, pair)
				return nil
			})
		}
		log.Printf("Finished parsing all values.\n")
		return nil
	})
	return resultPairs
}

func (repo RSSRepository) DeleteValue(key []byte) error {
	db := repo.OpenDatabase()
	defer db.Close()
	log.Printf("Deleting key-value pair in database with  key: %s\n", key)
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})
	if err != nil {
		log.Printf("Database encountered error: %s!\n", err)
		return err
	}
	return nil
}

func (repo RSSRepository) InsertValue(key []byte, value []byte) error {
	key = []byte("feeds:" + string(key))
	db := repo.OpenDatabase()
	defer db.Close()
	log.Printf("Inserting key-value pair in database: %s-%s\n", key, value)
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		if err != nil {
			log.Printf("Database encountered error: %s!\n", err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("Database encountered error: %s!\n", err)
		return err
	}
	return nil
}
