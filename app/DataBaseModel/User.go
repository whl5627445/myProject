package DataBaseModel

import (
	"gorm.io/gorm"
	"time"
)

type YssimUserSpace struct {
	ID            string         `gorm:"primaryKey"`
	SpaceName     string         `gorm:"column:space_name"`
	UserName      string         `gorm:"column:username"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     *time.Time     `gorm:"column:update_time"`
	LastLoginTime int64          `gorm:"column:last_login_time"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at"`
}
