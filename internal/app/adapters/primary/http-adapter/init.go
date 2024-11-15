package http_adapter

import (
	"context"
	"errors"
	"github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter/controller"
	controller_gen "github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter/controller-gen"
	"github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter/middleware"
	"github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter/router"
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	"log/slog"
	"net/http"

	"github.com/mkorobovv/full-restaurant/internal/app/config"
	"golang.org/x/sync/errgroup"
)

type HttpAdapter struct {
	server *http.Server
	config config.HttpAdapter
}

func New(config config.HttpAdapter, apiSvc *api_service.APIService) *HttpAdapter {
	r := router.New()

	ctr := controller.New(apiSvc)

	routerWithOptions := controller_gen.HandlerWithOptions(
		ctr,
		controller_gen.ChiServerOptions{
			BaseRouter: r.Router(),
			Middlewares: []controller_gen.MiddlewareFunc{
				middleware.CORS,
			},
		},
	)

	server := &http.Server{
		Handler:           routerWithOptions,
		Addr:              config.Port,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
	}

	return &HttpAdapter{
		server: server,
		config: config,
	}
}

func (a *HttpAdapter) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
		defer cancel()

		err := a.server.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		err := a.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	slog.Info("server started")

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}
