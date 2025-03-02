package dal

import (
	"github.com/cloudwego/gomall07/app/frontend/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
