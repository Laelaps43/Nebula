package system

type CasbinRule struct {
	ID    uint   `json:"id"`
	PType string `gorm:"column:ptype"`
	V0    string `gorm:"column:v0"`
	V1    string `gorm:"column:v1"`
	V2    string `gorm:"column:v2"`
	V3    string `gorm:"column:v3"`
	V4    string `gorm:"column:v4"`
	V5    string `gorm:"column:v5"`
}

// TableName 指定 CasbinRule 对应的表名
func (CasbinRule) TableName() string {
	return "casbin_rule"
}
