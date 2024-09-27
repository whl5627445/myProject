package config

import (
	"fmt"
	"time"

	"yssim-go/app/DataBaseModel"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openMySql() *gorm.DB {
	dsn := "root:root@tcp(mysql:3306)/yssim?charset=utf8mb4&parseTime=True&loc=Local"
	if DEBUG != "" {
		dsn = "root:root@tcp(" + DebugDB + ")/yssim?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败： %s", err))
	}
	SqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败： %s", err))
	}
	SqlDB.SetMaxIdleConns(30)
	SqlDB.SetMaxOpenConns(50)
	SqlDB.SetConnMaxLifetime(time.Hour * 1)
	Session := db.Session(&gorm.Session{PrepareStmt: true,
		Logger: logger.Default.LogMode(logger.Info)})
	return Session
}

var DB = openMySql()

func init() {
	DB.AutoMigrate(
		&DataBaseModel.YssimExperimentRecord{},
		&DataBaseModel.YssimSimulateRecord{},
		&DataBaseModel.YssimModelsCollection{},
		&DataBaseModel.YssimSnapshots{},
		&DataBaseModel.YssimUserSpace{},
		&DataBaseModel.YssimUserSettings{},
		&DataBaseModel.YssimModels{},
		&DataBaseModel.SystemLibrary{},
		&DataBaseModel.UserLibrary{},
		&DataBaseModel.ParameterCalibrationRecord{},
		&DataBaseModel.ParameterCalibrationTemplate{},

		&DataBaseModel.AppDataSource{},
		&DataBaseModel.AppSpace{},
		&DataBaseModel.AppPage{},
		&DataBaseModel.AppPageComponent{},
		&DataBaseModel.AppPageComponentsRelease{},
		&DataBaseModel.AppComponentBases{},

		&DataBaseModel.YssimResourceLib{},
	)
}
