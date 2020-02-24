package helper

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Database *gorm.DB

// 初始化连接
func init() {
	mysql := GetMysqlConfig()
	Database, err = gorm.Open(mysql.DB.Driver, mysql.DB.Username+":"+mysql.DB.Password+mysql.DB.Url)
	if err != nil {
		ErrLogger.Error("db connect error: " + err.Error())
		os.Exit(0)
	}
	Database.DB().SetMaxOpenConns(mysql.DB.MaxOpenConns)
	Database.DB().SetMaxIdleConns(mysql.DB.MaxIdleConns)
	Database.SetLogger(SQLLogger)
	Database.LogMode(mysql.DB.ShowSql)
	Database.SingularTable(mysql.DB.Singular)
}
