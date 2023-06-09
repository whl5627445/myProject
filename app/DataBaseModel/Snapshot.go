package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type YssimSnapshots struct {
	ID                string         `gorm:"primaryKey;type:varchar(128)"`
	SpaceId           string         `gorm:"column:space_id;type:varchar(128)"`
	PackageId         string         `gorm:"column:package_id;type:varchar(128)"`
	ExperimentId      string         `gorm:"column:experiment_id;type:varchar(128)"`
	SnapshotName      string         `gorm:"column:snapshot_name;type:varchar(32)"`
	UserName          string         `gorm:"column:username;type:varchar(32)"`
	ModelName         string         `gorm:"column:model_name;type:varchar(32)"`
	ComponentName     string         `gorm:"column:component_name;type:varchar(32)"`
	SimulateResultId  string         `gorm:"column:simulate_result_id;type:varchar(128)"`
	SimulateResultObj datatypes.JSON `gorm:"column:simulate_result_obj;type:json"`
	SimulateVarData   datatypes.JSON `gorm:"column:simulate_var_data;type:json"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data;type:json"`

	CreatedAt *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt *time.Time     `gorm:"column:update_time"`
	Deleted   gorm.DeletedAt `gorm:"column:deleted_at"`
}
