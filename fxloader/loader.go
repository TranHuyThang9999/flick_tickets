package fxloader

import (
	"flick_tickets/api/routers"
	repository "flick_tickets/core/adapter"

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
		//	fx.Provide(use_case.NewUseCaseUser),
	}
}

func loadValidator() []fx.Option {
	return []fx.Option{
		fx.Provide(validator.New),
	}
}
func loadEngine() []fx.Option {
	return []fx.Option{
		// fx.Provide(controllers.NewBaseController),
		// fx.Provide(controllers.NewUserController),
		// fx.Provide(controllers.NewAuthController),
		fx.Provide(routers.NewApiRouter),
		// fx.Provide(middlewares.NewMiddleware),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(repository.NewpostgreDb),
		//fx.Provide(postgres.NewpostgreDb),
		//fx.Provide(repositories.NewResposUser),
	}
}
