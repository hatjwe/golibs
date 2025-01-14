package logs

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger = initLogger()

func initLogger() *zap.Logger {
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig() //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	//encoderConfig := zap.NewDevelopmentEncoderConfig() //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder                      //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	if os.Getenv("debug") == "true" {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder //显示完整文件路径
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileName := "./server.log"
	//文件writeSyncer
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName, //日志文件存放目录
		MaxSize:    50,       //文件大小限制,单位MB
		MaxBackups: 5,        //最大保留日志文件数量
		MaxAge:     30,       //日志文件保留天数
		Compress:   false,    //是否压缩处理
	})
	var fileCore zapcore.Core
	if os.Getenv("debug") == "true" {
		fileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	} else {
		fileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.InfoLevel)
	}

	return zap.New(fileCore, zap.AddCaller()) //AddCaller()为显示文件名和行号
}
