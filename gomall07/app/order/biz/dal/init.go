package dal

import (
	"github.com/cloudwego/gomall07/app/order/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
