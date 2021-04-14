package storage

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const filesBucket = "files"

type Storage interface {
	StoreFile(key, data []byte) error
	GetFile(key []byte) ([]byte, error)
}

type dbstorage struct {
	db *bolt.DB
}

func (s *dbstorage) StoreFile(key, data []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists([]byte(filesBucket))
		if err != nil {
			return err
		}

		return bkt.Put(key, data)
	})
}

func (s *dbstorage) GetFile(key []byte) ([]byte, error) {
	var b []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(filesBucket))
		if bkt == nil {
			return fmt.Errorf("bucket %s does not exists", filesBucket)
		}

		b = bkt.Get(key)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return b, nil
}

func NewStorage(datadir string) (Storage, error) {
	db, err := bolt.Open(
		fmt.Sprintf("%s/.storage.db", datadir), 0600, &bolt.Options{Timeout: 10 * time.Second})

	if err != nil {
		return nil, err
	}

	return &dbstorage{
		db: db,
	}, nil
}
