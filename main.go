package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niloy104/simplebank/api"
	db "github.com/niloy104/simplebank/db/sqlc"
	"github.com/niloy104/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot create db connection pool: ", err)
	}
	defer pool.Close()

	store := db.NewStore(pool)

	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
