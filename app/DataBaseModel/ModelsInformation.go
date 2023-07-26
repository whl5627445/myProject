package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimModels struct {
	ID             string         `gorm:"primaryKey;type:varchar(128)"`
	PackageName    string         `gorm:"column:package_name;type:varchar(128)"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32)"`
	SysUser        string         `gorm:"column:sys_or_user;default:\"\";type:varchar(32)"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(256)"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt      *time.Time     `gorm:"column:update_time"`
	UserSpaceId    string         `gorm:"column:userspace_id;default:\"\";type:varchar(128)"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"master\";type:varchar(128)"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128)"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at"`
	Default        bool           `gorm:"column:default_version;default:0;type:bool"`
}

type SystemLibrary struct {
	ID             string         `gorm:"primaryKey;type:varchar(128)"`
	UserName       string         `gorm:"column:username;type:varchar(128)"`
	PackageName    string         `gorm:"column:package_name;type:varchar(128)"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32)"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(256)"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"master\";type:varchar(128)"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128)"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt      *time.Time     `gorm:"column:update_time"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at"`
}

type UserLibrary struct {
	ID             string         `gorm:"primaryKey;type:varchar(128)"`
	UserName       string         `gorm:"column:username;type:varchar(128)"`
	PackageName    string         `gorm:"column:package_name;type:varchar(128)"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32)"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(256)"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"master\";type:varchar(128)"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128)"`
	Used           bool           `gorm:"column:used;type:bool"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt      *time.Time     `gorm:"column:update_time"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at"`
}
