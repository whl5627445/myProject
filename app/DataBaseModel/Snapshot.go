package DataBaseModel

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type YssimSnapshots struct {
	ID                string         `gorm:"primaryKey"`
	SpaceId           string         `gorm:"column:space_id"`
	SnapshotName      string         `gorm:"column:snapshot_name"`
	UserName          string         `gorm:"column:username"`
	ModelName         string         `gorm:"column:model_name"`
	PackageId         string         `gorm:"column:package_id"`
	ComponentName     string         `gorm:"column:component_name"`
	ExperimentId      string         `gorm:"column:experiment_id"`
	SimulateResultId  string         `gorm:"column:simulate_result_id"`
	SimulateResultObj datatypes.JSON `gorm:"column:simulate_result_obj"`
	SimulateVarData   datatypes.JSON `gorm:"column:simulate_var_data"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data"`

	CreatedAt *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt *time.Time     `gorm:"column:update_time"`
	Deleted   gorm.DeletedAt `gorm:"column:deleted_at"`
}
