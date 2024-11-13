package restaurant_repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mkorobovv/full-restaurant/internal/app/config"
	"github.com/mkorobovv/full-restaurant/internal/pkg/repohelpers"
)

type RestaurantRepository struct {
	DB     *sql.DB
	config config.Database
}

func New(cfg config.Database) *RestaurantRepository {
	return &RestaurantRepository{
		config: cfg,
		DB:     nil,
	}
}

func (repo *RestaurantRepository) Start() {
	currentHostString := fmt.Sprintf("DB host: [%s:%s].", repo.config.Host, repo.config.Port)

	slog.Info(currentHostString + " Подключение...")

	connectionString := repohelpers.GetConnectionString(repo.config.Type, repo.config.Host, repo.config.Port, repo.config.User, repo.config.Password, repo.config.Name)

	db, err := sql.Open(repo.config.Type, connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	slog.Info(currentHostString + " Подключено!")

	repo.DB = db
}
