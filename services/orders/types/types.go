package types

import (
	"context"

	"github.com/Rishi5154/gRPC-order/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
