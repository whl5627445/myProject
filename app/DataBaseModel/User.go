package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type YssimUserSpace struct {
	ID                 string         `gorm:"primaryKey;type:varchar(128)"`
	SpaceName          string         `gorm:"column:space_name;type:varchar(128)"`
	UserName           string         `gorm:"column:username;type:varchar(32)"`
	Description        string         `gorm:"column:description;type:varchar(128)"`
	Background         string         `gorm:"column:background;type:varchar(128)"`
	IconColor          string         `gorm:"column:icon_color;type:varchar(128)"`
	Icon               string         `gorm:"column:icon;type:varchar(128)"`
	Collect            bool           `gorm:"column:collect;type:bool"`
	PackageInformation datatypes.JSON `gorm:"column:package_information;type:json"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt          *time.Time     `gorm:"column:update_time"`
	LastLoginTime      int64          `gorm:"column:last_login_time;type:int"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at"`
}

type YssimUserSettings struct {
	ID          string     `gorm:"primaryKey"`
	UserName    string     `gorm:"column:username;type:varchar(32)"`
	GridDisplay bool       `gorm:"column:grid_display;type:bool" json:"grid_display" binding:"required"`
	CreatedAt   *time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:update_time"`
}
