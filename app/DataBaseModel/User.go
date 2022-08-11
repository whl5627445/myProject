package DataBaseModel

import "time"

type YssimUserSpace struct {
	ID            string     `gorm:"primaryKey"`
	SpaceName     string     `gorm:"column:spacename"`
	UserName      string     `gorm:"column:username"`
	CreatedAt     *time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"column:update_time"`
	LastLoginTime *time.Time `gorm:"column:last_login_time"`
}
