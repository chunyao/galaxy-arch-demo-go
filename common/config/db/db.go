package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

// InitDbConfig 初始化Db
func InitDbConfig() {
	log.Info("初始化数据库 Mysql")
	var err error
	dsn := viper.GetString("db.dsn")
	maxIdleConns := viper.GetInt("db.maxIdleConns")
	maxOpenConns := viper.GetInt("db.maxOpenConns")
	connMaxLifetime := viper.GetInt("db.connMaxLifetime")
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	}); err != nil {
		panic(fmt.Errorf("初始化数据库失败: %s \n", err))
	}
	sqlDB, err := DB.DB()
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)                                    // 空闲连接数
		sqlDB.SetMaxOpenConns(maxOpenConns)                                    // 最大连接数
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifetime)) // 单位：秒
	}
	log.Info("Mysql: 数据库初始化完成")
}
