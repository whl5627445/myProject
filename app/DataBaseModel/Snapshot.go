package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type YssimSnapshots struct {
	ID                string         `gorm:"primaryKey;type:varchar(128);comment:视图快照的唯一标识"`
	SpaceId           string         `gorm:"column:space_id;type:varchar(128);comment:空间id"`
	PackageId         string         `gorm:"column:package_id;type:varchar(128);comment:包id"`
	ExperimentId      string         `gorm:"column:experiment_id;type:varchar(128);comment:实验id"`
	SnapshotName      string         `gorm:"column:snapshot_name;type:varchar(32);comment:快照名称"`
	UserName          string         `gorm:"column:username;type:varchar(32);comment:用户名"`
	ModelName         string         `gorm:"column:model_name;type:varchar(128);comment:模型名"`
	ComponentName     string         `gorm:"column:component_name;type:varchar(128);comment:组件名"`
	SimulateResultId  string         `gorm:"column:simulate_result_id;type:varchar(128);comment:实验结果id"`
	SimulateResultObj datatypes.JSON `gorm:"column:simulate_result_obj;type:json;comment:视图的仿真结果数据"`
	SimulateVarData   datatypes.JSON `gorm:"column:simulate_var_data;type:json;comment:仿真参数数据"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data;type:json;comment:模型参数数据"`

	CreatedAt *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt *time.Time     `gorm:"column:update_time"`
	Deleted   gorm.DeletedAt `gorm:"column:deleted_at"`
}
