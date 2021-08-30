package system

// Repository provides access to user repository.
type Repository interface {
	// AddUser saves a given user to the repository.
	AddUser(User) error
	// Repository provides access to the user and backup storage.
	// GetUser returns the user with given ID.
	GetUser(int) (User, error)
	// GetAllUsers returns all users saved in storage.
	GetAllUsers() []User
	// GetAllBackups returns a list of all backups for a given user ID.
	GetAllBackups(int) []Backup
	// Repository provides access to the backup storage.
	// AddBackup saves a given backup.
	AddBackup(Backup) error
}
