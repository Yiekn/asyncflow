package config

var Config YamlConfig

type YamlConfig struct {
	Mysql     MysqlConfig     `mapstructure:"mysql"`
	Snowflake SnowflakeConfig `mapstructure:"snowflake"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type SnowflakeConfig struct {
	WorkerId     int64 `mapstructure:"workerId"`
	DatacenterId int64 `mapstructure:"datacenterId"`
}

func InitConfig() {
	// TODO: 从 config.yaml 文件中读取配置
	mysqlConfig := MysqlConfig{
		Host:     "121.37.46.115",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Database: "asyncflow",
	}
	snowflakeConfig := SnowflakeConfig{
		WorkerId:     1,
		DatacenterId: 1,
	}

	Config = YamlConfig{
		Mysql:     mysqlConfig,
		Snowflake: snowflakeConfig,
	}
}
