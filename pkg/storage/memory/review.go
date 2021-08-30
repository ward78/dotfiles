package memory

import (
	"time"
)

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
