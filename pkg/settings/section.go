package settings

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MysqlSetting  `mapstructure:"mysql"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifetime"`
}
