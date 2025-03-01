package dal

import (
	"github.com/cloudwego/gomall07/app/user/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
