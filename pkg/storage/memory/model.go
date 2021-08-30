package memory

import (
	"time"
)

// User defines the properties of a user to be listed
type User struct {
	ID        int
	Name      string
	Brewery   string
	Abv       float32
	ShortDesc string
	Created   time.Time
}

// Backup defines a user backup
type Backup struct {
	ID        string
	UserID    int
	FirstName string
	LastName  string
	Score     int
	Text      string
	Created   time.Time
}
