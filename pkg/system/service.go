package system

type Service interface {
	// Service provides user adding operations.
	AddUser(...User)
	// Service provides user and backup listing operations.
	GetUser(int) (User, error)
	GetUsers() []User
	GetUserBackups(int) []Backup
	// Service provides backuping operations.
	AddUserBackup(Backup)
}
