package dbs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"
	"log"
	"os"
	"time"
)

var (
	mainDB *gorm.DB
)

func InitMysql() {
	dbUrl := "root:123456@Qwertyu@tcp(gz-cynosdbmysql-grp-fwjiy0vl.sql.tencentcdb.com:29021)/dat"
	url := fmt.Sprintf("%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", dbUrl)
	mainDB = NewMysql(url, 10)
}

func GetMainDB() *gorm.DB {
	return mainDB
}

func NewMysql(args string, maxCon int) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(fmt.Sprintf("Got errors when connect database, the errors is '%v'", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	idle := maxCon
	if maxCon/3 > 10 {
		idle = maxCon / 3
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(idle)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(maxCon)

	return db
}
