package fxloader

import (
	"flick_tickets/api/controllers"
	"flick_tickets/api/middlewares"
	"flick_tickets/api/routers"
	"flick_tickets/core/adapter"
	"flick_tickets/core/adapter/repository"
	"flick_tickets/core/events/caching"
	"flick_tickets/core/events/caching/cache"
	"flick_tickets/core/events/sockets"
	"flick_tickets/core/usecase"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadValidator()...),
		fx.Options(loadEngine()...),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(usecase.NewUseCaseUser),
		fx.Provide(usecase.NewUseCaseJwt),
		fx.Provide(usecase.NewUsecaseTicker),
		fx.Provide(usecase.NewUseCaseFile),
		fx.Provide(usecase.NewUsecaseOrder),
		fx.Provide(usecase.NewUseCaseAes),
		fx.Provide(usecase.NewUseCaseCustomer),
		fx.Provide(cache.NewCache),
		fx.Provide(sockets.NewManagerClient),
		fx.Provide(sockets.NewServer),
	}
}

func loadValidator() []fx.Option {
	return []fx.Option{
		fx.Provide(validator.New),
	}
}
func loadEngine() []fx.Option {
	return []fx.Option{

		fx.Provide(routers.NewApiRouter),
		fx.Provide(controllers.NewControllersUser),
		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewControllerAuth),
		fx.Provide(middlewares.NewMiddleware),
		fx.Provide(controllers.NewControllerTicket),
		fx.Provide(controllers.NewControllerFileLc),
		fx.Provide(controllers.NewControllerOrder),
		fx.Provide(controllers.NewControllerAes),
		fx.Provide(controllers.NewControllerCustomer),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(adapter.NewpostgreDb),
		fx.Provide(repository.NewCollectionUser),
		fx.Provide(repository.NewTransaction),
		fx.Provide(repository.NewConllectionFileStore),
		fx.Provide(repository.NewCollectionTickets),
		fx.Provide(repository.NewCollectionOrder),
		fx.Provide(caching.NewRedisDb),
		fx.Provide(repository.NewCollectionCustomer),
	}
}
