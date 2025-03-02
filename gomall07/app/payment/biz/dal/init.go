package dal

import (
	"github.com/cloudwego/gomall07/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
