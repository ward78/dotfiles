package memory

import (
	"fmt"
	"time"

	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/backuping"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/listing"
)

// Memory storage keeps data in memory
type Storage struct {
	users   []User
	backups []Backup
}

// Add saves the given user to the repository
func (m *Storage) AddUser(b adding.User) error {
	for _, e := range m.users {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newB := User{
		ID:        len(m.users) + 1,
		Created:   time.Now(),
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
	}
	m.users = append(m.users, newB)

	return nil
}

// Add saves the given backup in the repository
func (m *Storage) AddBackup(r backuping.Backup) error {
	found := false
	for b := range m.users {
		if m.users[b].ID == r.UserID {
			found = true
		}
	}

	if found {
		created := time.Now()
		id := fmt.Sprintf("%d_%s_%s_%d", r.UserID, r.FirstName, r.LastName, created.Unix())

		newR := Backup{
			ID:        id,
			Created:   created,
			UserID:    r.UserID,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Score:     r.Score,
			Text:      r.Text,
		}

		m.backups = append(m.backups, newR)
	} else {
		return listing.ErrNotFound
	}

	return nil
}

// Get returns a user with the specified ID
func (m *Storage) GetUser(id int) (listing.User, error) {
	var user listing.User

	for i := range m.users {

		if m.users[i].ID == id {
			user.ID = m.users[i].ID
			user.Name = m.users[i].Name
			user.Brewery = m.users[i].Brewery
			user.Abv = m.users[i].Abv
			user.ShortDesc = m.users[i].ShortDesc
			user.Created = m.users[i].Created

			return user, nil
		}
	}

	return user, listing.ErrNotFound
}

// GetAll return all users
func (m *Storage) GetAllUsers() []listing.User {
	var users []listing.User

	for i := range m.users {

		user := listing.User{
			ID:        m.users[i].ID,
			Name:      m.users[i].Name,
			Brewery:   m.users[i].Brewery,
			Abv:       m.users[i].Abv,
			ShortDesc: m.users[i].ShortDesc,
			Created:   m.users[i].Created,
		}

		users = append(users, user)
	}

	return users
}

// GetAll returns all backups for a given user
func (m *Storage) GetAllBackups(userID int) []listing.Backup {
	var list []listing.Backup

	for i := range m.backups {
		if m.backups[i].UserID == userID {
			r := listing.Backup{
				ID:        m.backups[i].ID,
				UserID:    m.backups[i].UserID,
				FirstName: m.backups[i].FirstName,
				LastName:  m.backups[i].LastName,
				Score:     m.backups[i].Score,
				Text:      m.backups[i].Text,
				Created:   m.backups[i].Created,
			}

			list = append(list, r)
		}
	}

	return list
}
