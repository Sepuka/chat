package cmd

import (
	"chat/src/def"
	"chat/src/def/source"
	commandSource "chat/src/source"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(terminalCmd)
}

var (
	terminalCmd = &cobra.Command{
		Use:     `terminal`,
		Example: `terminal -c /config/path [list|create]`,
		Short:   `hosting terminal controller`,
		Long:    `Manages hosting staff via terminal`,

		RunE: func(cmd *cobra.Command, args []string) error {
			commandSourceListener, err := def.Container.SafeGet(source.TerminalDef)
			if err != nil {
				return err
			}

			return commandSourceListener.(*commandSource.Terminal).Execute(args)
		},
	}
)
