package node

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

var (
	FilesBucket          = []byte("files")
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrCouldNotGetFile   = errors.New("file could not be found")
	ErrKeyAlreadyExists  = errors.New("key already exists")
)

type (
	File struct {
		ID   UID
		Name string
	}
)

// NewFile returns a new file struct
func (n *Node) NewFile(name string) (*File, error) {
	fullpath := filepath.Join(n.basepath, name)
	_, err := os.Stat(fullpath)

	if !os.IsNotExist(err) {
		return nil, ErrFileAlreadyExists
	}

	uid := NewFileUID(name)

	return &File{
		Name: name,
		ID:   uid,
	}, nil
}

// Store write the file and register in database
func (n *Node) Store(data []byte, file *File) error {
	fullpath := filepath.Join(n.basepath, file.Name)

	err := os.WriteFile(fullpath, data, os.ModePerm)
	if err != nil {
		return nil
	}

	err = n.database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(FilesBucket)
		if err != nil {
			return err
		}

		if k := b.Get(file.ID[:]); k != nil {
			return ErrKeyAlreadyExists
		}

		if err = b.Put(file.ID[:], []byte(file.Name)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		removeErr := os.Remove(fullpath)

		if removeErr != nil {
			return removeErr
		}

		return err
	}

	return nil
}

// Get return the file struct and the file data or error
func (n *Node) Get(fileID UID) (*File, []byte, error) {
	tx, err := n.database.Begin(false)
	if err != nil {
		return nil, nil, err
	}

	b := tx.Bucket(FilesBucket)
	filename := b.Get(fileID[:])

	if filename == nil {
		return nil, nil, ErrCouldNotGetFile
	}

	fullpath := filepath.Join(n.basepath, string(filename))

	filebytes, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return nil, nil, err
	}

	return &File{
		ID:   fileID,
		Name: string(filename),
	}, filebytes, nil
}
