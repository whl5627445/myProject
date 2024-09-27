package DataBaseModel

import (
	"time"

	"gorm.io/gorm"
)

type YssimResourceLib struct {
	ID          string         `gorm:"index;primaryKey;type:varchar(128);comment:资源文件夹文件唯一标识"`
	UserName    string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	ParentId    string         `gorm:"index;column:parent_id;type:varchar(128);comment:父节点id"`
	FolderFile  string         `gorm:"column:folder_file;type:varchar(32);comment:资源是文件夹还是文件"`
	FullPath    string         `gorm:"column:full_path;type:varchar(256);comment:资源文件夹文件节点路径"`
	Name        string         `gorm:"column:name;type:varchar(32);comment:资源文件夹文件名称"`
	Description string         `gorm:"column:description;type:varchar(256);comment:资源文件文件描述"`
	FilePath    string         `gorm:"column:file_path;type:varchar(256);comment:资源文件存放路径"`
	CreatedAt   *time.Time     `gorm:"column:create_time;autoCreateTime;comment:资源文件夹文件创建时间"`
	UpdatedAt   *time.Time     `gorm:"column:update_time;comment:资源文件夹文件更新时间"`
	Deleted     gorm.DeletedAt `gorm:"column:deleted_at;comment:资源文件夹文件删除时间"`
}
