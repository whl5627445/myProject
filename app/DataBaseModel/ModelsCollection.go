package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimModelsCollection struct {
	ID          string         `gorm:"index;column:id;primaryKey;type:varchar(128);comment:收藏模型的唯一识别标识"`
	PackageId   string         `gorm:"index;column:package_id;type:varchar(128);comment:模型所在package的唯一识别标识"`
	ModelName   string         `gorm:"column:model_name;type:varchar(128);comment:模型名称"`
	CreatedAt   *time.Time     `gorm:"column:create_time;type:datetime;autoCreateTime;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"column:update_time;type:datetime;comment:更新时间"`
	UserSpaceId string         `gorm:"column:userspace_id;default:\"0\";type:varchar(128);comment:收藏模型所在用户空间的唯一标识"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间"`
}
