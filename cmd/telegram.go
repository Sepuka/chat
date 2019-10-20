package cmd

import (
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/source"
	commandSource "github.com/sepuka/chat/internal/source"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(telegramCmd)
}

var (
	telegramCmd = &cobra.Command{
		Use:     `telegram`,
		Example: `telegram -c /config/path`,
		Short:   `hosting telegram controller`,
		Long:    `Manages hosting staff via telegram`,

		RunE: func(cmd *cobra.Command, args []string) error {
			commandSourceListener, err := def.Container.SafeGet(source.TelegramDef)
			if err != nil {
				return err
			}

			return commandSourceListener.(*commandSource.Telegram).Listen()
		},
	}
)
