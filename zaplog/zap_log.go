package zaplog

import (
	"os"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once sync.Once

	zapLogger     *zap.Logger
	sugaredLogger *zap.SugaredLogger
)

// InitSizeFile 初始 zap 日志，按大小切割和备份个数、文件有效期
// serviceName 服务名
// path        日志文件路径
// logLevel    日志级别 debug  info  error
// maxSize     每个日志文件保存大小 100M
// maxBackup   保留 N 个备份
// maxAge      保留 N 天
// callerSkip 提升的堆栈帧数，0=当前函数，1=上一层函数，....
func InitSizeFile(serviceName, path, logLevel string, maxSize, maxBackup, maxAge, callerSkip int) {

	once.Do(func() {
		// 日志分割
		hook := lumberjack.Logger{
			Filename:   path,      // 日志文件路径，默认 os.TempDir()
			MaxSize:    maxSize,   // 每个日志文件保存大小
			MaxBackups: maxBackup, // 保留N个备份，默认不限
			MaxAge:     maxAge,    // 保留N天，默认不限
			Compress:   true,      // 是否压缩，默认不压缩
		}

		// 打印到文件
		write := zapcore.AddSync(&hook)

		// 初始 zap 日志
		zapInit(write, serviceName, logLevel, callerSkip)
	})
}

// InitConsole 初始 zap 日志，输出到终端
// serviceName 服务名
// logLevel    日志级别 debug  info  error
// callerSkip 提升的堆栈帧数，0-当前函数，1-上一层函数，....
func InitConsole(serviceName, logLevel string, callerSkip int) {
	once.Do(func() {
		// 打印到控制台
		write := zapcore.AddSync(os.Stdout)

		// 初始 zap 日志
		zapInit(write, serviceName, logLevel, callerSkip)
	})
}

// zapInit 初始化 zap 基本信息
// write       文件描述符
// serviceName 服务名
// logLevel    日志级别 debug  info  error
// callerSkip 提升的堆栈帧数，0=当前函数，1=上一层函数，....
func zapInit(write zapcore.WriteSyncer, serviceName, logLevel string, callerSkip int) {

	// 设置日志级别
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "call",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapTimeEncoder,                 // 日志时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行消耗的时间转化成浮点型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器,格式化调用堆栈
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		write,
		level,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 提升打印的堆栈帧数
	addCallerSkip := zap.AddCallerSkip(callerSkip)
	// 开启文件及行号
	development := zap.Development()
	// 添加字段-服务器名称
	filed := zap.Fields(zap.String("serviceName", serviceName))

	// 构造日志
	zapLogger = zap.New(core, caller, addCallerSkip, development, filed)
	sugaredLogger = zapLogger.Sugar()
}

// zapTimeEncoder 日志时间格式
func zapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// IsInitNil 是否没有初始化
func IsInitNil() {
	if zapLogger == nil {
		InitConsole("goutil", "debug", 1)
	}
}

// Debug 调试
func Debug(msg string, args ...zap.Field) {
	IsInitNil()
	zapLogger.Debug(msg, args...)
}

// Debugf 调试
func Debugf(template string, args ...interface{}) {
	IsInitNil()
	sugaredLogger.Debugf(template, args...)
}

// Info 信息
func Info(msg string, args ...zap.Field) {
	IsInitNil()
	zapLogger.Info(msg, args...)
}

// Infof 信息
func Infof(template string, args ...interface{}) {
	IsInitNil()
	sugaredLogger.Infof(template, args...)
}

// Warn 警告
func Warn(msg string, args ...zap.Field) {
	IsInitNil()
	zapLogger.Warn(msg, args...)
}

// Warnf 警告
func Warnf(template string, args ...interface{}) {
	IsInitNil()
	sugaredLogger.Warnf(template, args...)
}

// Error 错误
func Error(msg string, args ...zap.Field) {
	IsInitNil()
	zapLogger.Error(msg, args...)
}

// Errorf 错误
func Errorf(template string, args ...interface{}) {
	IsInitNil()
	sugaredLogger.Errorf(template, args...)
}
