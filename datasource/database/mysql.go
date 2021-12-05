package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go_server/configs/db"
	"go_server/internel/data/model"
)

// NewMySqlDB 实例化数据库引擎方法：mysql的数据引擎
func NewMySqlDB(db_config db.DataBase) *gorm.DB {
	// 连接数据库
	dsn := db_config.User + ":" + db_config.Pwd + "@tcp(" + db_config.Host + ":" + db_config.Port + ")/" +
		db_config.Database + "?charset=" + db_config.Charset + "&parseTime=" + db_config.ParseTime + "&loc=" + db_config.Loc
	//     fmt.Println(dsn)
	var logLevel logger.LogLevel
	switch db_config.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	case "silent":
		logLevel = logger.Silent
	default:
		logLevel = logger.Info

	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true, // 自动忽略创建外键
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Duration(db_config.SlowThreshold) * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logLevel,                                                  // 日志级别
				IgnoreRecordNotFoundError: false,                                                     // 是否忽略记录器的 ErrRecordNotFound 错误
				Colorful:                  true,                                                      // 是否禁用颜色
			},
		),
	})
	if err != nil {
		fmt.Println(dsn)
		panic("failed to connect database")
	}

	// 根据实体创建表
	db.AutoMigrate(&model.PedesFlowInfo{})

	return db
}
