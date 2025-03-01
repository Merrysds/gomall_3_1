package dal

import (
	"github.com/cloudwego/gomall07/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
