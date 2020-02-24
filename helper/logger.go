package helper

import (
	"os"

	"github.com/leaderwolfpipi/logger"
)

var ErrLogger, AccessLogger, WorkLogger, SQLLogger *logger.RotateFileLogger

// 获取日志实例
func Logger(outPath string) *logger.RotateFileLogger {
	_, err := os.Stat(outPath)
	if os.IsNotExist(err) {
		// 不存在则创建
		os.Create(outPath)
	}

	// 返回日志实例
	return logger.NewRotateFileLogger(outPath)
}

// 初始化日志实例
func init() {
	WorkLogger = Logger("logs/work.log")
	AccessLogger = Logger("logs/access.log")
	ErrLogger = Logger("logs/error.log")
	SQLLogger = Logger("logs/sql.log")
}
