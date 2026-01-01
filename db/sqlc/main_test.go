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
var testDB DBTX

func TestMain(m *testing.M) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot create connection pool:", err)
	}

	testDB = pool
	testQueries = New(testDB)

	code := m.Run()

	pool.Close()
	os.Exit(code)
}
