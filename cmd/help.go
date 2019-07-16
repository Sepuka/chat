package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helpCmd)
}

var (
	helpCmd = &cobra.Command{
		Use:     `help`,
		Example: `help`,
		Short:   `command invokes help tutorial`,
		Long:    `
create [opts]
delete [opts]
`,
	}
)
