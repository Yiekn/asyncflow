package config

import (
	rlog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	config GlobalConfig
	once   sync.Once
)

type GlobalConfig struct {
	Mysql     MysqlConfig     `mapstructure:"mysql"`
	Snowflake SnowflakeConfig `mapstructure:"snowflake"`
	Log       LogConfig       `mapstructure:"log"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type SnowflakeConfig struct {
	WorkerId     int64 `mapstructure:"worker_id"`
	DatacenterId int64 `mapstructure:"datacenter_id"`
}

type LogConfig struct {
	LogPattern string `mapstructure:"log_pattern"`
	LogPath    string `mapstructure:"log_path"`
	SaveDays   uint   `mapstructure:"save_days"`
	Level      string `mapstructure:"level"`
}

func GetGlobalConfig() *GlobalConfig {
	once.Do(readConfig)
	return &config
}

func readConfig() {
	// 设置默认值
	setDefaultConfig()

	// 从YAML文件读取配置
	v := viper.New()
	v.SetConfigName("app")     // 配置文件名称(无扩展名)
	v.SetConfigType("yml")     // 配置文件类型
	v.AddConfigPath(".")       // 当前目录下的conf
	v.AddConfigPath("./conf")  // 相对路径
	v.AddConfigPath("../conf") // 上一级目录的conf

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Warnf("conf file not found, use default config: %+v", err)
		}
	} else {
		// 将配置文件的值绑定到结构体
		if err := v.Unmarshal(&config); err != nil {
			log.Errorf("conf parse failed: %+v", err)
		}
	}

	// 确保日志路径存在
	if config.Log.LogPath != "" && config.Log.LogPattern == "file" {
		ensureLogPath(config.Log.LogPath)
	}
}

func ensureLogPath(logPath string) {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err := os.MkdirAll(logPath, 0755)
		if err != nil {
			log.Fatalf("create log directory failed: %+v", err)
		}
	}
}

func setDefaultConfig() {
	config.Mysql = MysqlConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Database: "asyncflow",
	}

	config.Snowflake = SnowflakeConfig{
		WorkerId:     1,
		DatacenterId: 1,
	}

	config.Log = LogConfig{
		LogPattern: "stdout",
		LogPath:    "./logs",
		SaveDays:   7,
		Level:      "info",
	}
}

func InitConfig() {
	globalConfig := GetGlobalConfig()
	// 设置日志级别
	level, err := log.ParseLevel(globalConfig.Log.Level)
	if err != nil {
		panic("Invalid log level: " + globalConfig.Log.Level)
	}
	log.SetLevel(level)
	// 设置日志格式
	log.SetFormatter(&logFormatter{
		log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		}})
	// 打印文件位置，行号
	log.SetReportCaller(true)
	// 设置日志输出
	switch globalConfig.Log.LogPattern {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "file":
		writer, err := rlog.New(
			filepath.Join(globalConfig.Log.LogPath, "%Y%m%d%H%M.log"),
			rlog.WithLinkName(filepath.Join(globalConfig.Log.LogPath, "latest.log")), // 最新日志链接
			rlog.WithMaxAge(time.Duration(globalConfig.Log.SaveDays)*24*time.Hour),   // 保存天数
			rlog.WithRotationTime(24*time.Hour),                                      // 每天切割一次
		)
		if err != nil {
			panic("Failed to set file logger: " + err.Error())
		}
		log.SetOutput(writer)
	default:
		panic("Invalid log pattern: " + globalConfig.Log.LogPattern)
	}
}
