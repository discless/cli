package commands

import "github.com/spf13/cobra"

var NewCmd = &cobra.Command{
	Use: "new [bot | function | secret]",
	Short: "Create a new function, bot or secret with given name",
	Args: cobra.MinimumNArgs(1),
}
