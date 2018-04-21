package store

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/msgpack"
)

const (
	databaseFileName = "smarthut.db"
)

// DB represents store struct
type DB struct {
	*storm.DB
}

// NewStore initializes a new store
func NewStore(storepath string) (*DB, error) {
	// Use gob instead of JSON
	db, err := storm.Open(storepath+databaseFileName, storm.Codec(msgpack.Codec))
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func initDB() {
	// db, err := storm.Open(databaseFileName)
	// if err != nil {
	// 	log.Println(err)
	// }
}
