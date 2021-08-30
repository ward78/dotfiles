package system

type Service interface {
	// Service provides user system operations.
	AddUser(...User)
	// Service provides user and backup system operations.
	GetUser(int) (User, error)
	GetUsers() []User
	GetUserBackups(int) []Backup
	// Service provides system operations.
	AddUserBackup(Backup)
}
