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
		// 创建日志目录
		os.MkdirAll(outPath, 0755)
	}

	// 返回日志实例
	return logger.NewRotateFileLogger(outPath)
}

// 日志记录实例化与启动
func init_logger() {
	// 初始化记录器
	WorkLogger = Logger("logs/work")
	AccessLogger = Logger("logs/access")
	ErrLogger = Logger("logs/error")
	SQLLogger = Logger("logs/sql")

	// 启动记录器
	WorkLogger.Start()
	AccessLogger.Start()
	ErrLogger.Start()
	SQLLogger.Start()
}
