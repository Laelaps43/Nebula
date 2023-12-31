package model

import "time"

// NEBULA 每张表都有的数据结构
type NEBULA struct {
	CreatedAt time.Time `json:"-"` // 创建时间
	UpdatedAt time.Time `json:"-"` // 更新时间
}
