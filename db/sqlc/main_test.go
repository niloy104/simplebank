package db

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/niloy104/simplebank/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB DBTX

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot create connection pool:", err)
	}

	testDB = pool
	testQueries = New(testDB)

	code := m.Run()

	pool.Close()
	os.Exit(code)
}
