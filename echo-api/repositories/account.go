package repositories

import (
	"time"

	"gorm.io/gorm"
)

type (
	AccountRepository interface {
		Create(Account) (*Account, error)
		GetByUsername(string) (*Account, error)
	}

	Account struct {
		ID        string         `gorm:"column:id;primaryKey"`
		Username  string         `gorm:"column:username;unique"`
		Password  string         `gorm:"column:password"`
		CreatedAt time.Time      `gorm:"column:created_at"`       // Automatically managed by GORM for creation time
		UpdatedAt time.Time      `gorm:"column:updated_at"`       // Automatically managed by GORM for update time
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"` // Used for soft deletes (marking records as deleted without actually removing them from the database).
	}
)
