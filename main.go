package main

import (
	"context"
	database "test/src/database"

	"github.com/uptrace/bun"
)

func main() {
	database.ConnectDB()
	db, ctx := database.ConnectDB(), database.Ctx()

	// createalltable(ctx, db)
	// database.InsertData(db, ctx)
	// database.SelectData(db, ctx)
	// database.UpdateData(db, ctx)
	database.DeleteData(db, ctx)

}

func createalltable(ctx context.Context, db *bun.DB) {
	database.CreateTable(ctx, db)
}
