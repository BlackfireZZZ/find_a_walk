package handlers

type Handlers struct {
	AuthHandler  AuthHandler
	EventHandler EventHandler
	UserHandler  UserHandler
	TagsHandler  TagsHandler
}

func InitHandlers(
	userService UserService,
	eventService EventService,
	tagsService TagsService,
) *Handlers {

	return &Handlers{
		AuthHandler:  *NewAuthHandler(userService),
		EventHandler: *NewEventHandler(eventService),
		UserHandler:  *NewUserHandler(userService),
		TagsHandler:  *NewTagsHandler(tagsService),
	}
}
