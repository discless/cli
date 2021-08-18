package commands

import "github.com/spf13/cobra"

var NewCmd = &cobra.Command{
	Use: "new [bot | function]",
	Short: "Create a new function or bot with given name",
	Args: cobra.MinimumNArgs(2),
}
