package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimUserSpace struct {
	ID            string         `gorm:"primaryKey"`
	SpaceName     string         `gorm:"column:space_name"`
	UserName      string         `gorm:"column:username"`
	Description   string         `gorm:"column:description"`
	Background    string         `gorm:"column:background"`
	Icon          string         `gorm:"column:icon"`
	collect       bool           `gorm:"column:collect"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     *time.Time     `gorm:"column:update_time"`
	LastLoginTime int64          `gorm:"column:last_login_time"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at"`
}

type YssimUserSettings struct {
	//ID          string     `gorm:"AUTO_INCREMENT"`   坑，创建记录的时候自增失败
	UserName    string     `gorm:"column:username"`
	GridDisplay bool       `gorm:"column:grid_display" json:"grid_display" binding:"required"`
	CreatedAt   *time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:update_time"`
}
