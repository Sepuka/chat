package cmd

import (
	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/def"
	"github.com/sepuka/chat/internal/def/log"
	"github.com/sepuka/chat/internal/def/source"
	commandSource "github.com/sepuka/chat/internal/source"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	terminalCmd.Flags().StringVarP(&user, `user`, `u`, ``, `user's nickname'`)
	_ = terminalCmd.MarkFlagRequired(`user`)
	rootCmd.AddCommand(terminalCmd)
}

var (
	user        string
	terminalCmd = &cobra.Command{
		Use:     `terminal`,
		Example: `terminal instr=list -c /config/path -u vasya`,
		Short:   `hosting terminal controller`,
		Long:    `Manages hosting staff via terminal`,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			commandSourceListener, err := def.Container.SafeGet(source.TerminalDef)
			if err != nil {
				return err
			}

			logger := def.Container.Get(log.LoggerDef)
			req := context.NewRequest(user, args[0], args[1:]...)
			logger.(*zap.SugaredLogger).
				Info(
					`got terminal command`,
					zap.String(`user`, req.GetLogin()),
					zap.String(`command`, req.GetCommand()),
					zap.Strings(`args`, req.GetArgs()),
				)

			result, err := commandSourceListener.(*commandSource.Terminal).Execute(req)
			if err != nil {
				return err
			} else {
				logger.(*zap.SugaredLogger).
					Info(
						result.Msg,
						zap.String(`user`, req.GetLogin()),
						zap.String(`command`, req.GetCommand()),
						zap.Strings(`args`, req.GetArgs()),
					)
			}

			return nil
		},
	}
)
