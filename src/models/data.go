package models

import (
	"github.com/uptrace/bun"
)

type Data struct {
	bun.BaseModel `bun:"table:data,alias:u"`

	ID   int64  `bun:",pk,autoincrement"`
	Data string `json:"data"`
}
