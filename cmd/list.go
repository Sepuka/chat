package cmd

import (
	"github.com/sepuka/chat/order"
	"github.com/sepuka/chat/src/def"
	orderDef "github.com/sepuka/chat/src/def/order"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var (
	listCmd = &cobra.Command{
		Use:     `list`,
		Short:   `prints the list of servers`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var order order.Order
			if err := def.Container.Fill(orderDef.OrderListDef, &order); err != nil {
				return err
			}

			order.Run()

			return nil
		},
	}
)
