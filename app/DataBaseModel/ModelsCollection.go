package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimModelsCollection struct {
	ID          string         `gorm:"column:id;primaryKey;type:varchar(128)"`
	PackageId   string         `gorm:"column:package_id;type:varchar(128)"`
	ModelName   string         `gorm:"column:model_name;type:varchar(128)"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"column:update_time"`
	UserSpaceId string         `gorm:"column:userspace_id;default:\"0\";type:varchar(128)"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at"`
}
