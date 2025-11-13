package utils

import (
	"app/backend/common"
	"log"
	"path/filepath"

	"github.com/donnie4w/go-logger/logger"
)

var (
	Log = initLogger(filepath.Join(common.WordDir, common.LogPath))
)

// InitLogger pathFile: 日志全路径
func initLogger(path string) *logger.Logging {
	log.Println("初始化日志……" + path)
	l := logger.NewLogger()
	l.SetOption(&logger.Option{
		Level:     logger.LEVEL_INFO,
		Console:   true, // 控制台输出
		Format:    logger.FORMAT_LEVELFLAG | logger.FORMAT_SHORTFILENAME | logger.FORMAT_DATE | logger.FORMAT_MICROSECONDS,
		Formatter: "[{time}] {level} {file}: {message}\n",
		// size或者time模式
		FileOption: &logger.FileSizeMode{
			Filename:   path,
			Maxsize:    100 * 1024 * 1024,
			Maxbuckup:  1,
			IsCompress: true,
		},
	})

	return l
}
