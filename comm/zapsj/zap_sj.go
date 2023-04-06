package zapsj

import (
	"github.com/shijiu-xf/go-base/config"
	"github.com/shijiu-xf/go-base/constant/zero"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ZapSJ struct {
	err string
}

func (e *ZapSJ) Error() string {
	return e.err
}
func NewMyConfig(config config.LoggerConfig) zap.Config {
	var development bool
	var err error
	// 开发者模式默认开启
	if config.Development == zero.String {
		development = true
	}
	development, err = strconv.ParseBool(config.Development)
	if err != nil {
		panic(err)
	}
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(GetZapLogLevelCode(config.Level)),
		Development:      development,
		Encoding:         config.Encoding,
		EncoderConfig:    GetEncoderConfig(config.EncoderConfig),
		OutputPaths:      GetOutPaths(config.OutputPaths),
		ErrorOutputPaths: GetOutPaths(config.ErrorOutputPaths),
	}
}

// GetZapLogLevelCode 将配置文件中的日志级别转换为zap的码值，便于配置
func GetZapLogLevelCode(level string) zapcore.Level {
	// 默认为Debug级别
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "dpanic":
		zapLevel = zapcore.DPanicLevel
	case "panic":
		zapLevel = zapcore.PanicLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.DebugLevel
	}
	return zapLevel
}

// GetEncoderConfig 获取编码配置文件
func GetEncoderConfig(encoderConfig string) zapcore.EncoderConfig {
	if encoderConfig == "dev" {
		return zap.NewDevelopmentEncoderConfig()
	}
	if encoderConfig == "pro" {
		return zap.NewProductionEncoderConfig()
	}
	return zap.NewDevelopmentEncoderConfig()
}

func GetOutPaths(paths string) []string {
	outPaths := make([]string, 0)
	split := strings.Split(paths, ",")
	if paths == zero.String {
		outPaths = append(outPaths, "stderr")
		return outPaths
	}
	if len(split) == 0 {
		outPaths = append(outPaths, paths)
		return outPaths
	}
	for _, s := range split {
		outPaths = append(outPaths, s)
	}
	return outPaths
}
func getLogWriter(logFileCfg *config.LogFileConfig) zapcore.WriteSyncer {
	// 默认文件大小 500MB
	if logFileCfg.MaxSize <= 0 {
		logFileCfg.MaxSize = 500
	}
	syncL := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   logFileCfg.FileName,
			MaxSize:    logFileCfg.MaxSize,
			MaxAge:     logFileCfg.MaxAge,
			MaxBackups: logFileCfg.MaxBackups,
			LocalTime:  logFileCfg.LocalTime,
			Compress:   logFileCfg.Compress,
		},
	)
	// 并且打印到控制台
	syncSt := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncL, syncSt)
}

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()

	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。

	// 序列化时间。eg: 2022-09-01T19:11:35.921+0800
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// "time":"2022-09-01T19:11:35.921+0800"
	encodeConfig.TimeKey = "time"
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func InitZap(config config.LoggerConfig) {
	lg := log.Default()
	lg.Printf("[InitZap] 初始化")
	cfg := NewMyConfig(config)
	logger, err := cfg.Build()
	if err != nil {
		log.Panic(err)
	}
	zap.ReplaceGlobals(logger)
	lg.Printf("[日志配置] 初始化完成")
}

// InitSJZap 初始化一个分割日志文件的zap
func InitSJZap(c config.LogFileConfig) error {
	lg := log.Default()
	lg.Printf("[InitSJZap] 初始化")
	// 添加日志文件切割组件配置
	if reflect.DeepEqual(c, config.LoggerConfig{}) {
		return &ZapSJ{err: "[InitSJZap] 日志初始化失败，配置文件为空"}
	}
	lvl := new(zapcore.Level)
	err := lvl.UnmarshalText([]byte(c.Level))
	if err != nil {
		return err
	}
	core := zapcore.NewCore(getEncoder(), getLogWriter(&c), lvl)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	lg.Printf("[InitSJZap] 初始化完成")
	return nil
}

func ZapL() *zap.Logger {
	return zap.L()
}
func ZapS() *zap.SugaredLogger {
	return zap.S()
}
