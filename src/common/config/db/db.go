package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// db连接
var (
	MASTER = viper.GetStringMap("datasource.primary") // 默认主数据源
	DBs    = map[string]*RDBManager{}                 // 初始化时加载数据源到集合
)

// 连接管理器
type RDBManager struct {
	OpenTx bool     // 是否开启事务
	DsName string   // 数据源名称
	Db     *gorm.DB // 非事务实例
	Tx     *gorm.Tx // 事务实例
	Errors []error  // 操作过程中记录的错误
}

// 数据库配置
type DBConfig struct {
	Name string // 数据源名称
	Dsn  string // 地址连接
}

// InitDbConfig 初始化Db
func InitDbConfig() {
	log.Info("初始化数据库 Mysql")

	for k, v := range viper.GetStringMap("datasource.datasource") {
		log.Printf("初始化数据源 %s ", k)
		dsn, _ := v.(map[string]interface{})["dsn"]
		db, err := MysqlSetup(dsn.(string))
		if err != nil {
			log.Printf("数据库链接失败 %s ", err.Error())
			return
		}
		rdb := &RDBManager{
			Db:     db,
			DsName: k,
		}
		DBs[k] = rdb
	}

	log.Info("Mysql: 数据库初始化完成")
}
func MysqlSetup(dsn string) (*gorm.DB, error) {
	maxIdleConns := viper.GetInt("datasource.maxIdleConns")
	maxOpenConns := viper.GetInt("datasource.maxOpenConns")
	connMaxLifetime := viper.GetInt("datasource.connMaxLifetime")
	//启用打印日志

	newLogger := logger.New(
		log.StandardLogger(),
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level: Silent、Error、Warn、Info
			Colorful:      false,         // 禁用彩色打印
		},
	)

	if viper.GetString("datasource.sql.debug") == "true" {
		newLogger = logger.New(
			log.StandardLogger(),
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level: Silent、Error、Warn、Info
				Colorful:      false,       // 禁用彩色打印
			},
		)
	}
	dialector := mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})
	conn, err := gorm.Open(dialector, &gorm.Config{
		Logger:      newLogger,
		QueryFields: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})

	if err != nil {
		log.Print(err.Error())
		return conn, err
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Print("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(maxIdleConns) // 设置最大连接数
	sqlDB.SetMaxOpenConns(maxOpenConns) // 设置最大的空闲连接数
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifetime))
	return conn, nil
}
