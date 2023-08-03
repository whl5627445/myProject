package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimModels struct {
	ID             string         `gorm:"index;primaryKey;type:varchar(128);comment:package唯一标识"`
	LibraryId      string         `gorm:"index;column:library_id;type:varchar(128);comment:library库唯一标识"`
	PackageName    string         `gorm:"index;column:package_name;type:varchar(128);comment:package名称，一般称为包名或库的名字"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	SysUser        string         `gorm:"index;column:sys_or_user;default:\"\";type:varchar(32);comment:是用户模型的话则为用户名，系统模型则是sys"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(256);comment:用户模型所在路径，常驻系统模型则为空"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt      *time.Time     `gorm:"column:update_time;comment:更新时间"`
	UserSpaceId    string         `gorm:"column:userspace_id;default:\"\";type:varchar(128);comment:package所在用户空间的唯一识别标识"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool;comment:是否有版本控制"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"\";type:varchar(128);comment:版本控制分支"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128);comment:版本控制tag"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	Default        bool           `gorm:"column:default_version;default:0;type:bool;comment:是否默认加载"`
	Encryption     bool           `gorm:"column:encryption;default:0;type:bool;comment:是否是加密模型"`
}

type SystemLibrary struct {
	ID             string         `gorm:"index;primaryKey;type:varchar(128);comment:package唯一标识"`
	UserName       string         `gorm:"index;column:username;type:varchar(128);comment:用户名"`
	PackageName    string         `gorm:"index;column:package_name;type:varchar(128);comment:package名称，一般称为包名或库的名字"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(256);comment:package所在路径"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool;comment:是否有版本控制"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"\";type:varchar(128);comment:版本控制分支"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128);comment:版本控制tag"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt      *time.Time     `gorm:"column:update_time;comment:更新时间"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	Encryption     bool           `gorm:"column:encryption;default:0;type:bool;comment:是否是加密模型"`
}

type UserLibrary struct {
	ID                string         `gorm:"index;primaryKey;type:varchar(128);comment:package唯一标识"`
	UserName          string         `gorm:"index;column:username;type:varchar(128);comment:用户名"`
	PackageName       string         `gorm:"index;column:package_name;type:varchar(128);comment:package名称，一般称为包名或库的名字"`
	Version           string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	Used              bool           `gorm:"column:used;type:bool;comment:该package是否已经被某空间使用"`
	FilePath          string         `gorm:"column:file_path;default:\"\";type:varchar(256);comment:package所在路径"`
	VersionControl    bool           `gorm:"column:version_control;default:false;type:bool;comment:是否有版本控制"`
	VersionBranch     string         `gorm:"column:version_branch;default:\"master\";type:varchar(128);comment:版本控制分支"`
	VersionTag        string         `gorm:"column:version_tag;default:\"\";type:varchar(128);comment:版本控制tag"`
	RepositoryAddress string         `gorm:"column:repository_address;default:\"\";type:varchar(256);comment:git地址"`
	AnotherName       string         `gorm:"column:another_name;default:\"\";type:varchar(128);comment:别名"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt         *time.Time     `gorm:"column:update_time;comment:更新时间"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
}
