package prestigechain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const (
	userDbDir  = ".userDb"
	userDbFile = "user.db"
	userBucket = "userBucket"
)

type UserData struct {
	Username string
	Password string
}

type UserService struct {
	db *bolt.DB
}

func NewUserService() (*UserService, error) {
	// initialize user db
	fullDbPath := userDbDir + "/" + userDbFile
	if _, err := os.Stat(fullDbPath); os.IsNotExist(err) {
		err = os.MkdirAll(userDbDir, 0700)
		if err != nil {
			return nil, err
		}
	}

	db, err := bolt.Open(fullDbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &UserService{
		db: db,
	}, nil
}

func (us *UserService) AddNewUser(username, password string) error {
	err := us.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(userBucket))
		encodedUsername := encodeString(username)
		encodedPassword := encodeString(password)

		if b.Get(encodedUsername) != nil {
			return fmt.Errorf("User exists")
		}

		err := b.Put(encodedUsername, encodedPassword)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (us *UserService) UserExists(username string) bool {
	err := us.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(userBucket))
		encodedUsername := encodeString(username)

		if b.Get(encodedUsername) == nil {
			return fmt.Errorf("User does not exist")
		}

		return nil
	})

	return err == nil
}

func (us *UserService) ValidateUserPassword(username, password string) error {
	err := us.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(userBucket))
		encodedUsername := encodeString(username)
		encodedPassword := encodeString(password)

		storedEncodedPassword := b.Get(encodedUsername)

		if storedEncodedPassword == nil {
			return fmt.Errorf("User does not exist")
		}

		if !bytes.Equal(encodedPassword, storedEncodedPassword) {
			return fmt.Errorf("Password is invalid")
		}

		return nil
	})

	return err
}

func (us *UserService) Delete() {
	us.db.Close()

	if _, err := os.Stat(userDbDir); !os.IsNotExist(err) {
		err = os.RemoveAll(userDbDir)
		if err != nil {
			log.Panic(err)
		}
	}
}

func encodeString(s string) []byte {
	encoded := sha256.Sum256([]byte(s))
	return encoded[:]
}
