package store

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/marceldegraaf/smartmeter/types"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbPath     = "/var/lib/smartmeter/usage.db"
	bucketName = "Usage"
)

var db *bolt.DB

func Initialize() {
	createDatabase()
	createBucket()

	log.Infof("Storage initialized, saving to: %s", dbPath)
}

func Save(usage types.Usage) {
	log.Debugf("Storing usage: %v", usage)

	err := save(usage)
	if err != nil {
		log.Errorf("Could not save usage: %s", err)
	}
}

func save(usage types.Usage) error {
	key := []byte(usage.Timestamp.Format(time.RFC3339))

	data, err := bson.Marshal(usage)
	if err != nil {
		log.Errorf("Could not marshal %#v to bson: %s", usage, err)
		return err
	}

	log.Debugf("Writing to key %q: %#v", key, data)

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		err := bucket.Put(key, data)
		if err != nil {
			log.Errorf("Could not store usage: %s", err)
			return err
		}

		return nil
	})

	log.Debugf("Stored usage data for key %q", key)

	return nil
}

func createDatabase() {
	var err error

	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}
}

func createBucket() {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatalf("Could not create bucket: %s", err)
			return err
		}

		return nil
	})
}
