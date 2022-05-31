package database

import (
	"context"
	"fmt"
	models "test/src/models"
	"time"

	bun "github.com/uptrace/bun"
)

func InsertData(db *bun.DB, ctx context.Context) {
	test := &models.TestTable{Name: "test", Data: "test"}
	_, err := db.NewInsert().Model(test).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func SelectData(db *bun.DB, ctx context.Context) {
	test := new(models.TestTable)
	err := db.NewSelect().ColumnExpr("name").ColumnExpr("data").Model(test).Where("id = ?", 1).Scan(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(*test)
}

func UpdateData(db *bun.DB, ctx context.Context) {
	test := &models.TestTable{Name: "changed", Data: "changed", CreatedAt: time.Now()}
	_, err := db.NewUpdate().Model(test).Where("id=?", 1).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func DeleteData(db *bun.DB, ctx context.Context) {
	test := &models.TestTable{}
	_, err := db.NewDelete().Model(test).Where("id = ?", 1).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
