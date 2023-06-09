package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimModels struct {
	ID          string         `gorm:"primaryKey;type:varchar(128)"`
	PackageName string         `gorm:"column:package_name;type:varchar(32)"`
	Version     string         `gorm:"column:version;default:\"\";type:varchar(32)"`
	SysUser     string         `gorm:"column:sys_or_user;default:\"\";type:varchar(32)"`
	FilePath    string         `gorm:"column:file_path;default:\"\";type:varchar(128)"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"column:update_time"`
	UserSpaceId string         `gorm:"column:userspace_id;default:\"1000\";type:varchar(128)"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at"`
	Default     bool           `gorm:"column:default_version;default:0;type:bool"`
}
