package helper

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var Database *gorm.DB

// 初始化连接
func init() {
	// 载入配置
	err := LoadConfig()

	// 判断错误
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 未找到配置文件
			WorkLogger.Info("no such config file!")
		} else {
			// 其他类型错误
			ErrLogger.Fatal("read config error: " + err.Error())
		}
	}

	// 实例化db
	mysql := GetMysqlConfig()
	Database, err = gorm.Open(mysql.Mysql.Driver, mysql.Mysql.Username+":"+mysql.Mysql.Password+mysql.Mysql.Url)
	if err != nil {
		ErrLogger.Error("db connect error: " + err.Error())
		os.Exit(0)
	}
	Database.DB().SetMaxOpenConns(mysql.Mysql.MaxOpenConns)
	Database.DB().SetMaxIdleConns(mysql.Mysql.MaxIdleConns)
	Database.SetLogger(SQLLogger)
	// Database.LogMode(mysql.DB.ShowSql)
	Database.SingularTable(mysql.Mysql.Singular)
}
