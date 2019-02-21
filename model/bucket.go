package model

import (
	"errors"
	"time"

	"github.com/smarthut/smarthut/store"
)

var (
	// ErrBucketNoName is returned if a new bucket has no name
	ErrBucketNoName = errors.New("unable to create a bucket without a name")
)

// JSONBucket holds a data with JSON data
type JSONBucket struct {
	ID   int                    `json:"id" storm:"id,increment"` // bucket ID
	Name string                 `json:"name" storm:"unique"`     // unique bucket name
	Data map[string]interface{} `json:"data"`                    // bucket data

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBucket creates a new bucket with a data
func NewBucket(name string, data map[string]interface{}) (*JSONBucket, error) {
	if name == "" {
		return nil, ErrBucketNoName
	}
	return &JSONBucket{
		Name:      name,
		Data:      data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// AllBuckets returns all available buckets
func AllBuckets(db *store.DB) ([]JSONBucket, error) {
	var buckets []JSONBucket
	if err := db.All(&buckets); err != nil {
		return nil, err
	}
	return buckets, nil
}

// GetBucket finds JSON bucket by it's name
func GetBucket(db *store.DB, name string) (*JSONBucket, error) {
	var bucket JSONBucket
	if err := db.One("Name", name, &bucket); err != nil {
		return nil, err
	}
	return &bucket, nil
}

// UpdateData updated data in JSON bucket
func (j *JSONBucket) UpdateData(db *store.DB, data map[string]interface{}) error {
	j.Data = data
	j.UpdatedAt = time.Now()
	return db.Update(j)
}

// Delete deletes JSON bucket from a database
func (j *JSONBucket) Delete(db *store.DB) error {
	return db.DeleteStruct(j)
}
