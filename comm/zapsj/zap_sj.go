package zapsj

import (
	"github.com/shijiu-xf/go-base/config"
	"github.com/shijiu-xf/go-base/constant/zero"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strconv"
	"strings"
)

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

func InitZap(config config.LoggerConfig) {
	lg := log.Default()
	lg.Printf("[日志配置] 初始化")
	cfg := NewMyConfig(config)
	logger, err := cfg.Build()
	if err != nil {
		log.Panic(err)
	}
	zap.ReplaceGlobals(logger)
	lg.Printf("[日志配置] 初始化完成")
}

func ZapL() *zap.Logger {
	return zap.L()
}
func ZapS() *zap.SugaredLogger {
	return zap.S()
}
