package model

var (
	_cfg *config
)

func GetConfig() *config {
	if nil == _cfg {
		_cfg = new(config)
	}
	return _cfg
}

type config struct {
	Log      logConfig                 `toml:"log"`    // log output path
	Server   serverConfig              `toml:"server"` // server log level
	Database map[string]databaseConfig `toml:"database"`
	// redis配置
	Redis RedisConfig
}

type logConfig struct {
	Dir      string `toml:"dir"`       // 存储目录路径
	Category string `toml:"category"`  // 日志分类目录
	Level    string `toml:"level"`     // 日志级别
	StdPrint bool   `toml:"std_print"` // 是否打印到控制台
}

// 服务配置
type serverConfig struct {
	Addr string `toml:"addr"` // 服务地址
}

// 数据库配置
type databaseConfig struct {
	Driver          string            // 数据库驱动
	Source          string            // 数据库源
	ConnMaxLifeTime int               // 数据库最大连接时长
	Log             databaseLogConfig `toml:"log"` // 数据库日志
}

// 数据库日志配置
type databaseLogConfig struct {
	Mode     int    // 日志模式： 0-无日志, 1-写日志, 2-读写日志
	Category string // 日志类别
}

type RedisConfig struct {
	MedTest []*RedisConnect `toml:"med_test"`
}

type RedisConnect struct {
	Addr           string `toml:"addr"` // 地址
	Auth           string `toml:"auth"` // 密码
	Db             int    `toml:"db"`   // 数据库
	Idle           int    `toml:"idle"` // 最大连接数
	Active         int    // 一次性活跃
	Wait           bool   // 是否等待空闲连接
	ConnectTimeout int64  `toml:"connect_timeout"` // 连接超时时间， 毫秒
}
