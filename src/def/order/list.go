package order

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/chat/order"
	"github.com/sepuka/chat/src/def"
)

const OrderListDef = "order.list"

func init() {
	def.Register(func(builder *di.Builder, cfg def.Config) error {
		return builder.Add(di.Def{
			Name: OrderListDef,
			Tags: []di.Tag{{
				Name: order.OrderDefTag,
			}},
			Build: func(ctx def.Context) (interface{}, error) {
				return &order.List{}, nil
			},
		})
	})
}
