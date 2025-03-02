package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/gomall07/app/payment/biz/model"
	"github.com/cloudwego/gomall07/app/payment/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dns := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.PaymentLog{})
	if err != nil {
		panic(err)
	}
}
