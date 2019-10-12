package commands

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/MichaelMure/git-bug/bridge/core"
)

var (
	bridgeToken core.Token
)

func runBridgeTokenAdd(cmd *cobra.Command, args []string) error {
	if err := bridgeToken.Validate(); err != nil {
		return errors.Wrap(err, "invalid token")
	}

	if bridgeToken.Global {
		return core.StoreToken(repo, &bridgeToken)
	}

	return core.StoreGlobalToken(repo, &bridgeToken)
}

var bridgeTokenAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Configure and use bridge tokens.",
	PreRunE: loadRepo,
	RunE:    runBridgeTokenAdd,
	Args:    cobra.NoArgs,
}

func init() {
	bridgeTokenCmd.AddCommand(bridgeTokenAddCmd)
	bridgeTokenAddCmd.Flags().BoolVarP(&bridgeToken.Global, "global", "g", false, "")
	bridgeTokenAddCmd.Flags().StringVarP(&bridgeToken.Value, "value", "v", "", "")
	bridgeTokenAddCmd.Flags().StringVarP(&bridgeToken.Target, "target", "t", "", "")
	bridgeTokenAddCmd.Flags().StringArrayVarP(&bridgeToken.Scopes, "scopes", "s", []string{}, "")
	bridgeTokenAddCmd.Flags().SortFlags = false
}
