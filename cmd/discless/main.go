package main

import (
	"github.com/discless/cli/discless/commands"
	"github.com/spf13/cobra"
)

func main() {
	// Args
	commands.INewBot()
	commands.INewSecret()
	commands.IUp()

	// Commands
	var rootCmd = &cobra.Command{Use: "discless"}
	rootCmd.AddCommand(commands.NewCmd)
	rootCmd.AddCommand(commands.DeployCmd)
	rootCmd.AddCommand(commands.UpCmd)
	rootCmd.AddCommand(commands.StartCmd, commands.StopCmd)
	commands.NewCmd.AddCommand(commands.NewFuncCmd, commands.NewBotCmd, commands.NewSecretCmd)
	rootCmd.Execute()
}