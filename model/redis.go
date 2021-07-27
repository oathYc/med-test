package model

import "github.com/gomodule/redigo/redis"

var RedisPool *redis.Pool

type RedisConnectConfig struct {
	Addr           string `toml:"addr"`            // 地址
	Auth           string `toml:"auth"`            // 密码
	Db             int    `toml:"db"`              // 数据库
	Idle           int    `toml:"idle"`            // 最大连接数
	Active         int    `toml:"active"`          // 一次性活跃
	Wait           bool   `toml:"wait"`            // 是否等待空闲连接
	ConnectTimeout int64  `toml:"connect_timeout"` // 连接超时时间， 毫秒
}
