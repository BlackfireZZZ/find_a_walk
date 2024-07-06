package services

import (
	"github.com/go-chi/jwtauth/v5"
)

type Services struct {
	UserService  UserService
	EventService EventService
	TagService   TagService
}

func InitServices(
	userRepo UserRepository, jwtConfig *jwtauth.JWTAuth,
	eventRepo EventRepository, tagRepo TagRepository,
) *Services {

	return &Services{
		UserService:  *NewDefaultUserService(userRepo, jwtConfig),
		EventService: *NewDefaultEventService(eventRepo),
		TagService:   *NewDefaultTagService(tagRepo),
	}
}
