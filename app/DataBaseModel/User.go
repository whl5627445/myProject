package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type YssimUserSpace struct {
	ID                 string         `gorm:"primaryKey;type:varchar(128);comment:用户空间唯一标识"`
	SpaceName          string         `gorm:"column:space_name;type:varchar(128);comment:用户空间名称"`
	UserName           string         `gorm:"column:username;type:varchar(32);comment:用户名"`
	Description        string         `gorm:"column:description;type:varchar(128);comment:空间描述简介"`
	Background         string         `gorm:"column:background;type:varchar(128);comment:空间背景图所在目录"`
	IconColor          string         `gorm:"column:icon_color;type:varchar(128);comment:空间图标颜色"`
	Icon               string         `gorm:"column:icon;type:varchar(128);comment:空间图标"`
	Collect            bool           `gorm:"column:collect;type:bool;comment:是否被收藏"`
	PackageInformation datatypes.JSON `gorm:"column:package_information;type:json;comment:表示当前空间已加载哪些库和版本，以及所在目录"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt          *time.Time     `gorm:"column:update_time;comment:更新时间"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	LastLoginTime      int64          `gorm:"column:last_login_time;type:int;comment:最近进入该空间的时间戳"`
}

type YssimUserSettings struct {
	ID          string     `gorm:"primaryKey;comment:用户配置的唯一标识"`
	UserName    string     `gorm:"column:username;type:varchar(32);comment:用户名"`
	GridDisplay bool       `gorm:"column:grid_display;type:bool;comment:栅格是否显示"`
	CreatedAt   *time.Time `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt   *time.Time `gorm:"column:update_time;comment:更新时间"`
}
