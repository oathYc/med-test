package boot

import (
	"hello/model"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func InitRedisPool() error {
	if model.GetConfig() == nil || 0 == len(model.GetConfig().Redis.MedTest) {
		return errors.New("redis配置有误")
	}
	redisConfig := model.GetConfig().Redis.MedTest[0]
	// 尝试是否可以连接
	dial := func() (redis.Conn, error) {
		return redis.Dial(
			"tcp",
			redisConfig.Addr,
			redis.DialPassword(redisConfig.Auth),
			redis.DialDatabase(redisConfig.Db),
			redis.DialConnectTimeout(time.Duration(redisConfig.ConnectTimeout)*time.Millisecond),
		)
	}
	conn, err := dial()
	if err != nil {
		return errors.WithMessage(err, "初始化redis连接失败")
	}
	// 只是验证是否能够连接
	err = conn.Close()
	if err != nil {
		return errors.WithMessage(err, "初始化redis连接失败")
	}
	model.RedisPool = &redis.Pool{
		Dial:      dial,
		MaxIdle:   redisConfig.Idle,
		MaxActive: redisConfig.Active,
		Wait:      redisConfig.Wait,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
