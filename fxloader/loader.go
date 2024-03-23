package fxloader

import (
	"flick_tickets/api/controllers"
	"flick_tickets/api/middlewares"
	"flick_tickets/api/routers"
	"flick_tickets/core/adapter"
	"flick_tickets/core/adapter/repository"
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
		fx.Provide(usecase.NewAesUseCase),
		fx.Provide(usecase.NewJwtUseCase),
		fx.Provide(usecase.NewUsecaseTicker),
		fx.Provide(usecase.NewUseCaseFile),
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
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(adapter.NewpostgreDb),
		fx.Provide(repository.NewCollectionUser),
		fx.Provide(repository.NewTransaction),
		fx.Provide(repository.NewConllectionFileStore),
		fx.Provide(repository.NewCollectionTickets),
	}
}
