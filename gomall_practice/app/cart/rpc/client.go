package rpc

import (
	"sync"

	"github.com/cloudwego/gomall07/app/cart/conf"
	cartUtils "github.com/cloudwego/gomall07/app/cart/utils"
	clientsuite "github.com/cloudwego/gomall07/common/clientsuite"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
