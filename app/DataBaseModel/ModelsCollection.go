package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimModelsCollection struct {
	ID          string         `gorm:"column:id;primaryKey"`
	PackageId   string         `gorm:"column:package_id"`
	ModelName   string         `gorm:"column:model_name"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"column:update_time"`
	UserSpaceId string         `gorm:"column:userspace_id;default:\"0\""`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at"`
}
