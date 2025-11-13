package service

import (
	"app/backend/common"
	"app/backend/utils"
	"path/filepath"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	DBFilePath string
)

func init() {
	NewDB()
}

func NewDB() {
	st := time.Now()
	defer func() {
		utils.Log.Infof("初始化数据库完成，耗时 %s", time.Since(st))
	}()

	appConfig := NewAppConfig()
	c := appConfig.GetConfig()

	// sqlite3
	DBFilePath = filepath.Join(c.DBDir, common.LocalDBFile)
	//mode=rwc确保文件可读写并创建。
	//_journal_mode=WAL提升容错性和读写性能。
	dsn := DBFilePath + "?mode=rwc&_journal_mode=WAL"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置自动 VACUUM。
	//auto_vacuum = FULL会在删除数据时自动回收空间，但需在数据库初始化时设置，对现有数据库无效。
	db.Exec("PRAGMA auto_vacuum = FULL;")
	// SQLite 是文件数据库，通常只需要一个连接
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxOpenConns(1) // 最大打开连接数为 1
	sqlDB.SetMaxIdleConns(1) // 不保留空闲连接，任务完成后立即关闭
	DB = db
}
