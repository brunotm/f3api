package badgerdb

/*
Copyright 2018 Bruno Moura <brunotm@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"bytes"
	"os"

	"github.com/dgraph-io/badger"
)

// Store type
type Store struct {
	path string
	db   *badger.DB
}

// Open or creates a store
func Open(path string) (store *Store, err error) {
	if err = os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	opts.NumVersionsToKeep = 1
	opts.MaxTableSize = 8 << 20      // def 64 << 20
	opts.NumMemtables = 3            // def 5
	opts.NumLevelZeroTables = 3      // def 5
	opts.NumLevelZeroTablesStall = 5 // def 10

	var db *badger.DB
	db, err = badger.Open(opts)
	if err != nil {
		return store, err
	}

	store = &Store{}
	store.db = db
	store.path = path

	return store, nil

}

func (s *Store) Close() (err error) {
	return s.db.Close()
}

func (s *Store) Get(key []byte, callback func(value []byte, err error)) {
	s.db.View(
		func(txn *badger.Txn) error {
			item, err := txn.Get(key)
			if err == badger.ErrKeyNotFound {
				callback(nil, nil)
				return nil
			}

			if err != nil {
				callback(nil, err)
				return nil
			}

			value, err := item.Value()
			if err != nil {
				callback(nil, err)
				return nil
			}

			callback(value, nil)
			return nil

		},
	)
}

func (s *Store) Set(key, value []byte) (err error) {
	return s.db.Update(
		func(txn *badger.Txn) error {
			return txn.Set(key, value)
		},
	)
}

func (s *Store) Delete(key []byte) (err error) {
	return s.db.Update(
		func(txn *badger.Txn) error {
			return txn.Delete(key)
		},
	)
}

func (s *Store) Iter(from, to []byte, callback func(key, value []byte) error) (err error) {
	return s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10

		it := txn.NewIterator(opts)
		defer it.Close()

		it.Rewind()

		if len(from) > 0 {
			it.Seek(from)
		}

		for {

			if !it.Valid() {
				return nil
			}

			item := it.Item()
			key := item.Key()

			if len(to) > 0 {
				if bytes.Compare(key, to) > 0 {
					return nil
				}
			}

			value, err := item.Value()
			if err != nil {
				return err
			}

			if err = callback(key, value); err != nil {
				return err
			}

			it.Next()
		}

	})
}
