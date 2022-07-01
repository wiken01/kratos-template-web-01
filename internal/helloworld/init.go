package helloworld

import (
	"os"

	"github.com/natefinch/lumberjack"
	"web_template/pkg/log"
)

type LogWriteType int

const (
	LogWriteTypeFile LogWriteType = iota
	LogWriteTypeOsStd
)

func LogWriter(writeType ...LogWriteType) log.WriteSyncer {
	if len(writeType) == 0 {
		return os.Stdout
	}

	w := make([]log.WriteSyncer, 0)
	for _, v := range writeType {
		switch v {
		case LogWriteTypeFile:
			w = append(w, &lumberjack.Logger{
				Filename:   "log.log", // 日志文件路径
				MaxSize:    100,       // 每个日志文件保存的大小 单位:M
				MaxAge:     30,        // 文件最多保存10天
				MaxBackups: 100,       // 日志文件最多保存多少个备份
				Compress:   false,     // 是否压缩
			})
		case LogWriteTypeOsStd:
			w = append(w, os.Stdout)
		}
	}
	return log.NewMultiWriteSyncer(w...)
}
