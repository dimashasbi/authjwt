package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// SysAdmin is model for database configuration system
type SysAdmin struct {
	Name  string         `gorm:"size:40;index:key_registry;unique"`
	Value postgres.Jsonb `gorm:"size:2000"`
}

// NewSysAdmin create new SysAdmin
func NewSysAdmin(name, jsonmessage string) *SysAdmin {
	return &SysAdmin{
		Name: name,
		Value: postgres.Jsonb{
			RawMessage: json.RawMessage(jsonmessage),
		},
	}
}
