package database

import (
	"fmt"
	"gin-server/src/utils/config"
	"gin-server/src/utils/mylogs"
	"github.com/tiantianlikeu/gorm-plus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

// 全局变量 ,初始化后的连接
var Orm *orm

type orm struct {
	myDb *gorm.DB
}

func NewOrm() *orm {
	o := new(orm)
	o.initDBConfig()
	// gorm 增强版
	gorm_plus.Init(o.myDb)
	return o
}

func init() {
	Orm = NewOrm()
}

func (db *orm) initDBConfig() (err error) {

	if db.myDb != nil {
		return
	}

	password := config.Config.Db.Password
	username := config.Config.Db.Username
	host := config.Config.Db.Host
	port := config.Config.Db.Port
	database := config.Config.Db.Database
	charset := config.Config.Db.Charset
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True"
	dsn = fmt.Sprintf(dsn, username, password, host, port, database, charset)

	// 初始化GORM日志配置
	f, logfileErr := mylogs.GetLogFile()
	var logfile *os.File
	if logfileErr == nil {
		logfile = f
	} else {
		logfile = os.Stdout
	}
	writer := io.MultiWriter(logfile, os.Stdout)
	newLogger := logger.New(
		log.New(writer, "\r\n", log.Ldate|log.Ltime|log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		fmt.Println("数据库连接错误：%v", err)
		return err
	}

	// 尝试与数据库建立连接
	sqlDB, err := gormDb.DB()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	db.myDb = gormDb

	fmt.Println("数据库初始化成功")
	return nil
}

func (db *orm) DB() *gorm.DB {
	return db.myDb
}
