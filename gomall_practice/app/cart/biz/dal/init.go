package dal

import (
	"github.com/cloudwego/gomall07/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
