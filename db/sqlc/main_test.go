package db

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot create connection pool:", err)
	}
	defer pool.Close()

	testQueries = New(pool)
	os.Exit(m.Run())
}
