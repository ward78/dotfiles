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
