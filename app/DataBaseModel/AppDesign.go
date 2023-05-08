package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type AppSpace struct {
	ID            string         `gorm:"primaryKey"`
	SpaceName     string         `gorm:"column:space_name"`
	UserName      string         `gorm:"column:username"`
	Description   string         `gorm:"column:description"`
	Background    string         `gorm:"column:background"`
	Icon          string         `gorm:"column:icon"`
	IconColor     string         `gorm:"column:icon_color"`
	Collect       bool           `gorm:"column:collect"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     *time.Time     `gorm:"column:update_time"`
	LastLoginTime int64          `gorm:"column:last_login_time"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppDataSource struct {
	ID            string         `gorm:"primaryKey"`
	UserName      string         `gorm:"column:username"`
	UserSpaceId   string         `gorm:"column:user_space_id"`
	PackageId     string         `gorm:"column:package_id"`
	ModelName     string         `gorm:"column:model_name"`
	CompilerType  string         `gorm:"column:compiler_type"`
	CompilePath   string         `gorm:"column:compile_path"`
	CompileStatus int            `gorm:"column:compile_status"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppPage struct {
	ID           string         `gorm:"primaryKey"`
	UserName     string         `gorm:"column:username"`
	AppSpaceId   string         `gorm:"column:app_space_id"`
	PageName     string         `gorm:"column:page_name"`
	PagePath     string         `gorm:"column:page_path"`
	DataSourceId string         `gorm:"column:data_source_id"`
	PageWidth    string         `gorm:"column:page_width"`
	PageHeight   string         `gorm:"column:page_height"`
	Background   string         `gorm:"column:background"`
	Color        string         `gorm:"column:color"`
	Release      bool           `gorm:"column:is_release"`
	CreatedAt    *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt    *time.Time     `gorm:"column:update_time"`
	Deleted      gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppComponent struct {
	ID                 string         `gorm:"primaryKey"`
	PageId             string         `gorm:"column:page_id"`
	DataObjectList     datatypes.JSON `gorm:"column:data_object_list"`
	Type               string         `gorm:"column:type"`
	Width              string         `gorm:"column:width"`
	Height             string         `gorm:"column:height"`
	PositionX          float64        `gorm:"column:position_x"`
	PositionY          float64        `gorm:"column:position_y"`
	Angle              float64        `gorm:"column:angle"`
	HorizontalFlip     bool           `gorm:"column:horizontal_flip"`
	VerticalFlip       bool           `gorm:"column:vertical_flip"`
	Opacity            int            `gorm:"column:opacity"`
	OtherConfiguration datatypes.JSON `gorm:"column:other_configuration"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt          *time.Time     `gorm:"column:update_time"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at"`
}
