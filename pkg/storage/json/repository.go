package json

import (
	"encoding/json"
	"fmt"

	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/backuping"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/listing"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionUser identifier for the JSON collection of users
	CollectionUser = "users"
	// CollectionBackup identifier for the JSON collection of backups
	CollectionBackup = "backups"
)

// Storage stores user data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddUser saves the given user to the repository
func (s *Storage) AddUser(b adding.User) error {

	existingUsers := s.GetAllUsers()
	for _, e := range existingUsers {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newB := User{
		ID:        len(existingUsers) + 1,
		Created:   time.Now(),
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
	}

	resource := strconv.Itoa(newB.ID)
	if err := s.db.Write(CollectionUser, resource, newB); err != nil {
		return err
	}
	return nil
}

// AddBackup saves the given backup in the repository
func (s *Storage) AddBackup(r backuping.Backup) error {

	var user User
	if err := s.db.Read(CollectionUser, strconv.Itoa(r.UserID), &user); err != nil {
		return listing.ErrNotFound
	}

	created := time.Now()
	newR := Backup{
		ID:        fmt.Sprintf("%d_%s_%s_%d", r.UserID, r.FirstName, r.LastName, created.Unix()),
		Created:   created,
		UserID:    r.UserID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
	}

	if err := s.db.Write(CollectionBackup, newR.ID, r); err != nil {
		return err
	}

	return nil
}

// Get returns a user with the specified ID
func (s *Storage) GetUser(id int) (listing.User, error) {
	var b User
	var user listing.User

	var resource = strconv.Itoa(id)

	if err := s.db.Read(CollectionUser, resource, &b); err != nil {
		// err handling omitted for simplicity
		return user, listing.ErrNotFound
	}

	user.ID = b.ID
	user.Name = b.Name
	user.Brewery = b.Brewery
	user.Abv = b.Abv
	user.ShortDesc = b.ShortDesc
	user.Created = b.Created

	return user, nil
}

// GetAll returns all users
func (s *Storage) GetAllUsers() []listing.User {
	list := []listing.User{}

	records, err := s.db.ReadAll(CollectionUser)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var b User
		var user listing.User

		if err := json.Unmarshal([]byte(r), &b); err != nil {
			// err handling omitted for simplicity
			return list
		}

		user.ID = b.ID
		user.Name = b.Name
		user.Brewery = b.Brewery
		user.Abv = b.Abv
		user.ShortDesc = b.ShortDesc
		user.Created = b.Created

		list = append(list, user)
	}

	return list
}

// GetAll returns all backups for a given user
func (s *Storage) GetAllBackups(userID int) []listing.Backup {
	list := []listing.Backup{}

	records, err := s.db.ReadAll(CollectionBackup)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, b := range records {
		var r Backup

		if err := json.Unmarshal([]byte(b), &r); err != nil {
			// err handling omitted for simplicity
			return list
		}

		if r.UserID == userID {
			var backup listing.Backup

			backup.ID = r.ID
			backup.UserID = r.UserID
			backup.FirstName = r.FirstName
			backup.LastName = r.LastName
			backup.Score = r.Score
			backup.Text = r.Text
			backup.Created = r.Created

			list = append(list, backup)
		}
	}

	return list
}
