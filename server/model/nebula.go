package model

import "time"

// 每张表都有的数据结构
type NEBULA struct {
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
