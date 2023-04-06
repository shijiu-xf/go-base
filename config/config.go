package config

import "github.com/asim/go-micro/v3/config"

type AllConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	SrvInfo    SrvConfig    `mapstructure:"server" json:"server"`
	MySQLInfo  MySQLConfig  `mapstructure:"mysql" json:"mysql"`
	LoggerInfo LoggerConfig `mapstructure:"logger" json:"logger"`
	JwtInfo    JwtInfo      `mapstructure:"jwt" json:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}

type ConsulConfig struct {
	Host                string              `mapstructure:"host" json:"host"`
	Port                string              `mapstructure:"port" json:"port"`
	ConsulAddr          string              `json:"consulAddr"`
	ConfigurationCenter ConfigurationCenter `mapstructure:"config-center" json:"config-center"`
}
type ConfigurationCenter struct {
	Prefix string `mapstructure:"prefix" json:"prefix"`
}

type SrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Schema   string `mapstructure:"schema" json:"schema"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

// LoggerConfig struct 的配置字段含义请参考 zap.Config
type LoggerConfig struct {
	Level            string `mapstructure:"level" json:"level"`
	Development      string `mapstructure:"development" json:"development"`
	Encoding         string `mapstructure:"encoding" json:"encoding"`
	EncoderConfig    string `mapstructure:"encoderConfig" json:"encoderConfig"`
	OutputPaths      string `mapstructure:"outputPaths" json:"outputPaths"`
	ErrorOutputPaths string `mapstructure:"errorOutputPaths" json:"errorOutputPaths"`
	// 日志文件配置
	//LogFileConfig LogFileConfig `mapstructure:"logFileConfig" json:"logFileConfig"`
}

// LogFileConfig 日志文件配置，属性意义请参考 lumberjack.Logger
type LogFileConfig struct {
	Level      string `mapstructure:"level" json:"level"`
	FileName   string `mapstructure:"fileName" json:"fileName"`
	MaxSize    int    `mapstructure:"maxSize" json:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge" json:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups" json:"maxBackups"`
	LocalTime  bool   `mapstructure:"localTime" json:"localTime"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
	// 是否开启日志监控
	LogMonitorEnable bool `mapstructure:"logMonitorEnable" json:"logMonitorEnable"`
}

// JwtInfo 字段意义参考 jwt.RegisteredClaims
type JwtInfo struct {
	Key      string   `mapstructure:"key" json:"key"`
	Expires  int64    `mapstructure:"expires" json:"expires"`
	Issuer   string   `mapstructure:"issuer" json:"issuer"`
	Subject  string   `mapstructure:"subject" json:"subject"`
	Audience []string `mapstructure:"audience" json:"audience"`
	Platform string   `mapstructure:"platform" json:"platform"`
}

// JaegerConfig 链路追踪配置
type JaegerConfig struct {
	Addr        string `json:"addr"`
	ServiceName string `json:"serviceName"`
}

func (c *AllConfig) GetServerConfig(config config.Config, path ...string) error {
	err := config.Get(path...).Scan(c)
	return err
}
func (c *MySQLConfig) GetMysqlConfig(config config.Config, path ...string) error {
	err := config.Get(path...).Scan(c)
	return err
}
func (c *SrvConfig) GetUserSrvConfig(config config.Config, path ...string) error {
	err := config.Get(path...).Scan(c)
	return err
}
func (c *LoggerConfig) GetLoggerConfig(config config.Config, path ...string) error {
	err := config.Get(path...).Scan(c)
	return err
}
func (c *JwtInfo) GetJwtInfo(config config.Config, path ...string) error {
	err := config.Get(path...).Scan(c)
	return err
}
func (c *JaegerConfig) GetJaegerConfig(config config.Config, path ...string) error {
	err := config.Get(path...).Scan(c)
	return err
}
