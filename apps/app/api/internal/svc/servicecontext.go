package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shopping/apps/app/api/internal/config"
	"shopping/apps/order/rpc/order"
	"shopping/apps/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
