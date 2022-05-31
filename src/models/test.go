package models

import (
	"time"

	"github.com/uptrace/bun"
)

type TestTable struct {
	bun.BaseModel `bun:"table:test_table,alias:u"`

	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",notnull" json:"name"`
	Data      string    `bun:",notnull" json:"data"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:now()"`
}
