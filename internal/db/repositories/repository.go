package repositories

import "database/sql"

type Repositories struct {
	UserRepository *UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
