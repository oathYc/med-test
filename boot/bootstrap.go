package boot

import (
	"git.medlinker.com/golang/xerror"
)

func Bootstrap(configPath string) (err error) {

	if err = bootConfig(configPath); nil != err {
		return xerror.Wrap(err, "bootConfig")
	}

	// 初始化日志
	if err = bootLogger(); err != nil {
		return xerror.Wrap(err, "bootLogger")
	}

	if err = bootMysql(); nil != err {
		return xerror.Wrap(err, "bootMysql")
	}

	if err = InitRedisPool(); nil != err {
		return xerror.Wrap(err, "bootRedis")
	}
	return nil
}
