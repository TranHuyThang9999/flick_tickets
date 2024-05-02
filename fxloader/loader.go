package fxloader

import (
	"flick_tickets/api/controllers"
	"flick_tickets/api/middlewares"
	"flick_tickets/api/public/assets/infrastructure"
	"flick_tickets/api/public/assets/services"
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
		fx.Provide(usecase.NewUseCaseShowTime),
		fx.Provide(usecase.NewUseCaseCinemas),
		fx.Provide(services.NewServiceAddress),
		fx.Provide(usecase.NewUseCasePayment),
		fx.Provide(usecase.NewUseCaseMovie),
		fx.Provide(usecase.NewUseCaseCart),
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
		fx.Provide(controllers.NewControllerShowTIme),
		fx.Provide(controllers.NewControllerCinamas),
		fx.Provide(controllers.NewControllerAddress),
		fx.Provide(controllers.NewControllerParment),
		fx.Provide(controllers.NewControllerMovie),
		fx.Provide(controllers.NewControllerCart),
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
		fx.Provide(repository.NewCollectionShowTime),
		fx.Provide(repository.NewCollectionCinemas),
		fx.Provide(infrastructure.NewCollectionAddress),
		fx.Provide(repository.NewCollectionMovie),
		fx.Provide(repository.NewCollectionCart),
	}
}
