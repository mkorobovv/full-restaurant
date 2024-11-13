package main

import (
	"context"
	restaurant_repository "github.com/mkorobovv/full-restaurant/internal/app/adapters/secondary/repositories/restaurant-repository"
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	"log"
	"log/slog"
	"os"

	http_adapter "github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter"
	"github.com/mkorobovv/full-restaurant/internal/pkg/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	config := config.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	repo := restaurant_repository.New(config.Adapters.Secondary.Databases.Restaurant)

	repo.Start()

	apiSvc := api_service.New(repo)

	httpAdapter := http_adapter.New(config.Adapters.Primary.HttpAdapter, apiSvc)

	g.Go(func() error {
		err := httpAdapter.Start(ctx)
		if err != nil {
			log.Println(err)

			log.Println("Start graceful shutdown")

			return err
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		slog.Error(err.Error())

		os.Exit(1)
	}
}
