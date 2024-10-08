package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

// 管网信息文件表
type YssimPipeNetCad struct {
	ID          string         `gorm:"index;primaryKey;type:varchar(128);comment:管网信息文件记录唯一标识"`
	UserName    string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	Name        string         `gorm:"column:name;type:varchar(32);comment:管网信息文件记录名称"`
	Description string         `gorm:"column:description;type:varchar(256);comment:管网信息文件描述"`
	Path        string         `gorm:"column:path;type:varchar(256);comment:管网信息文件存放路径"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime;comment:管网信息文件记录创建时间"`
	UpdatedAt   *time.Time     `gorm:"column:update_time;comment:管网信息文件记录更新时间"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at;comment:管网信息文件记录删除时间"`
}

// 映射配置表
type YssimMappingConfig struct {
	ID          string         `gorm:"index;primaryKey;type:varchar(128);comment:映射配置记录唯一标识"`
	UserName    string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	Name        string         `gorm:"column:name;type:varchar(32);comment:映射配置记录名称"`
	Description string         `gorm:"column:description;type:varchar(256);comment:映射配置描述"`
	Path        string         `gorm:"column:path;type:varchar(256);comment:映射配置文件存放路径"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime;comment:映射配置记录创建时间"`
	UpdatedAt   *time.Time     `gorm:"column:update_time;comment:映射配置记录更新时间"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at;comment:映射配置记录删除时间"`
}

// 实例映射表
type YssimInstanceMapping struct {
	ID              string         `gorm:"index;primaryKey;type:varchar(128);comment:实例映射记录唯一标识"`
	UserName        string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	PipeNetCadId    string         `gorm:"column:pipe_net_cad_id;type:varchar(128);comment:管网信息文件记录id"`
	MappingConfigId string         `gorm:"column:mapping_config_id;type:varchar(128);comment:映射配置记录id"`
	Path            string         `gorm:"column:path;type:varchar(256);comment:实例映射文件存放路径"`
	PackageId       string         `gorm:"column:package_id;type:varchar(128);comment:模型id"`
	CreatedAt       *time.Time     `gorm:"column:create_time;autoCreateTime;comment:实例映射记录创建时间"`
	UpdatedAt       *time.Time     `gorm:"column:update_time;comment:实例映射记录更新时间"`
	Deleted         gorm.DeletedAt `gorm:"column:deleted_at;comment:实例映射记录删除时间"`
}

// 管网信息文件下载表
type YssimPipeNetCadDownload struct {
	ID          string         `gorm:"index;primaryKey;type:varchar(128);comment:管网信息文件记录唯一标识"`
	UserName    string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	Name        string         `gorm:"column:name;type:varchar(32);comment:管网信息文件记录名称"`
	Description string         `gorm:"column:description;type:varchar(256);comment:管网信息文件描述"`
	PipeNetPath string         `gorm:"column:pipe_net_path;type:varchar(256);comment:管网信息文件存放路径"`
	MappingPath string         `gorm:"column:mapping_path;type:varchar(256);comment:映射表存放路径"`
	PackageId   string         `gorm:"column:package_id;type:varchar(128);comment:模型id"`
	ModelName   string         `gorm:"column:model_name;type:varchar(128);comment:模型名"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime;comment:管网信息文件记录创建时间"`
	UpdatedAt   *time.Time     `gorm:"column:update_time;comment:管网信息文件记录更新时间"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at;comment:管网信息文件记录删除时间"`
}
