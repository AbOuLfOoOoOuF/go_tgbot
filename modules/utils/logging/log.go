package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sysSeparator = string(os.PathSeparator) // either \ or /

func InitLogger() *zap.Logger {
	var s = sysSeparator
	path, _ := os.Getwd()
	if _, err := os.Stat(path + s + "logs"); os.IsNotExist(err) {
		os.Mkdir(path+s+"logs", os.ModePerm)
	}
	if _, err := os.Stat(path + s + "logs" + s + "logs.txt"); os.IsNotExist(err) {
		os.Create(path + s + "logs" + s + "logs.txt")
	}

	file, err := os.OpenFile(path+s+"logs"+s+"logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	PanicErr(err)
	// wrSync := zapcore.AddSync(file)
	wrSync := zap.CombineWriteSyncers(file, zapcore.Lock(os.Stdout))

	encConf := zap.NewProductionEncoderConfig()
	encConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encConf.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encConf)

	core := zapcore.NewCore(encoder, wrSync, zap.InfoLevel)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(zapLogger)

	zapLogger.Sync()
	return zapLogger
}

// Panic on errors
func PanicErr(err error) {
	if err != nil {
		zap.S().Panic(err.Error())
		zap.S().Panic(err)
	}
}

func Panic(err string) {
	zap.S().Panic(err)
}

// handle errors
func HandleErr(err error) {
	if err != nil {
		zap.S().Error(err.Error())
		zap.S().Error(err)
	}
}

func Info(msg string) {
	zap.S().Info(msg)
}
