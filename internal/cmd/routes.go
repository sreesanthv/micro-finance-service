package cmd

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sreesanthv/micro-finance-service/internal/controller"
	"github.com/sreesanthv/micro-finance-service/internal/domain"
	"github.com/sreesanthv/micro-finance-service/internal/repo"
	"github.com/sreesanthv/micro-finance-service/internal/service"
)

func Routes(db *pgxpool.Pool, redis domain.Redis) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	accountRepo := repo.NewAccontRepo(db, logger)
	accountService := service.NewAccountService(accountRepo, logger)

	accountController := controller.NewAccontController(accountService, logger, redis)
	appMiddlewares := controller.NewMiddleware(accountController)

	r.Route("/account", func(r chi.Router) {
		r.Post("/", accountController.Create)
		r.Post("/login", accountController.Login)

		r.Group(func(r chi.Router) {
			r.Use(appMiddlewares.Auth)
			r.Post("/logout", accountController.Logout)

			r.Get("/{id:[0-9]+}", accountController.Get)
			r.Put("/{id:[0-9]+}", accountController.Update)
			r.Delete("/{id:[0-9]+}", accountController.Delete)

			r.Post("/debit", accountController.Debit)
			r.Post("/credit", accountController.Credit)
		})
	})

	return r
}
