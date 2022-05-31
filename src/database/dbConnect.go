package database

import (
	"context"
	"database/sql"
	models "test/src/models"

	_ "github.com/lib/pq"
	bun "github.com/uptrace/bun"
	pgdialect "github.com/uptrace/bun/dialect/pgdialect"
	pgdriver "github.com/uptrace/bun/driver/pgdriver"
	bundebug "github.com/uptrace/bun/extra/bundebug"
)

func Ctx() context.Context {
	ctx := context.Background()
	return ctx
}

func ConnectDB() *bun.DB {

	// Open a PostgreSQL database.
	dsn := "postgres://postgres:mysecretpassword@localhost:5432/test2?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db
}

func CreateTable(ctx context.Context, db *bun.DB) {

	// resp1, err := db.NewCreateTable().Model((*models.TestTable)(nil)).Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(resp1)
	// }
	// resp2, err := db.NewCreateTable().Model((*models.Data)(nil)).Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(&resp2)
	// }
	err := db.ResetModel(ctx, (*models.Data)(nil), (*models.TestTable)(nil))
	if err != nil {
		panic(err)
	}
}
