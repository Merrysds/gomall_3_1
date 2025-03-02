package rpc

import (
	"sync"

	"github.com/cloudwego/gomall07/app/frontend/conf"
	frontendUtils "github.com/cloudwego/gomall07/app/frontend/utils"
	frontendutils "github.com/cloudwego/gomall07/app/frontend/utils"
	"github.com/cloudwego/gomall07/common/clientsuite"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = frontendutils.ServiceName
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func Init() {
	once.Do(func() {
		iniUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func iniUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleErr(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleErr(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleErr(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleErr(err)
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleErr(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleErr(err)
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleErr(err)
}
