package dal

import (
	"github.com/cloudwego/gomall04/biz/dal/mysql"
	"github.com/cloudwego/gomall04/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
