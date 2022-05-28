package initialize

import (
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

const (
	InfoLogPath = "./logs/info.log"
	ErrLogPath  = "./logs/error.log"
)

func Zap() *zap.SugaredLogger {
	// Zap 配置
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "ts",
		CallerKey:    "file",
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}

	// 自定义日志级别: INFO
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= zap.InfoLevel
	})
	// 自定义日志级别: Warn
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= zap.InfoLevel
	})

	core := zapcore.NewTee(
		// 将 INFO 及以下级别的日志输出到文件
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(writer(InfoLogPath)), infoLevel),
		// 将 WARN 及以上级别的日志输出到文件
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(writer(ErrLogPath)), warnLevel),
		// 将日志输出到控制台
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zap.InfoLevel),
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)).Sugar()
}

func writer(filename string) io.Writer {
	hook, err := rotateLogs.New(
		filename+".%Y-%m-%d-%H",
		rotateLogs.WithLinkName(filename),
		rotateLogs.WithMaxAge(time.Hour*24*7),
		rotateLogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(fmt.Errorf("RotateLogs error: %v", err))
	}
	return hook
}
