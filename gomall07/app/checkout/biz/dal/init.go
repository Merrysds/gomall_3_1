package dal

import (
	"github.com/cloudwego/gomall07/app/checkout/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
