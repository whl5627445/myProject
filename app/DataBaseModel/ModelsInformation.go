package DataBaseModel

import "time"

type YssimModels struct {
	ID          string `gorm:"primaryKey"`
	PackageName string `gorm:"column:package_name"`
	//ModelName   string     `gorm:"column:model_name"`
	SysUser     string     `gorm:"column:sys_or_user"`
	FilePath    string     `gorm:"column:file_path"`
	CreatedAt   *time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:update_time"`
	UserSpaceId string     `gorm:"column:userspace_id;default:\"0\""`
}