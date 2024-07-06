package repositories

import "github.com/jackc/pgx/v4/pgxpool"


type Repositories struct {
	EventRepository EventRepository
	TagRepository   TagRepository
	UserRepository  UserRepository
}

func InitRepositores(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		EventRepository: *NewEventRepository(db),
		TagRepository:   *NewTagRepository(db),
		UserRepository:  *NewUserRepository(db),
	}
}
