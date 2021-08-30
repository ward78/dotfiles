package system

import "errors"

// ErrNotFound is used when a user could not be found.
var ErrNotFound = errors.New("user not found")
var ErrDuplicate = errors.New("user already exists")

const (
	// Done means finished processing successfully
	Done Event = iota
	// UserAlreadyExists means the given user is a duplicate of an existing one
	UserAlreadyExists
	// Failed means processing did not finish successfully
	Failed
	// We could also have a Queued Event which would mean queued for processing
	Queued
)

// Event defines possible outcomes from the "system actor"
type Event int

func (e Event) GetMeaning() string {
	if e == Done {
		return "Done"
	}

	if e == UserAlreadyExists {
		return "Duplicate user"
	}

	if e == Failed {
		return "Failed"
	}

	return "Unknown result"
}

type service struct {
	r Repository
}

// NewService creates an system service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser adds the given user(s) to the database
func (s *service) AddUser(b ...User) {
	// any validation can be done here
	for _, user := range b {
		_ = s.r.AddUser(user) // error handling omitted for simplicity
	}
}

// GetUsers returns all users
func (s *service) GetUsers() []User {
	return s.r.GetAllUsers()
}

// GetUser returns a user
func (s *service) GetUser(id int) (User, error) {
	return s.r.GetUser(id)
}

// GetUserBackups returns all requests for a user
func (s *service) GetUserBackups(userID int) []Backup {
	return s.r.GetAllBackups(userID)
}

// AddUserBackup saves a new user backup in the database
func (s *service) AddUserBackup(r Backup) {
	_ = s.r.AddBackup(r) // error handling omitted for simplicity
}
