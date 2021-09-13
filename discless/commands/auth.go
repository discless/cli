package commands

import (
	"fmt"
	"github.com/discless/cli/discless"
	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth [bot] [key]",
	Short: "add a key to the key register to be manage your bot through the API",
	Args:  cobra.MinimumNArgs(2),
	RunE:  AuthF,
}

func AuthF(c *cobra.Command, args []string) error {
	if err := discless.AddKey(args[0], args[1]); err != nil {
		return err
	}
	fmt.Println("✅ Added key to key register")
	return nil
}
