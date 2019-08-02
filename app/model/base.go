package model

import (
	"github.com/gogf/gf/g/os/gtime"
)

type Base struct {
	ID        uint `gconv:"id"`
	CreatedAt gtime.Time
	UpdatedAt gtime.Time
}
