package main

import (
	"github.com/discless/discless-cli/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "discless"}
	rootCmd.AddCommand(commands.NewCmd)
	rootCmd.AddCommand(commands.ApplyCmd)
	commands.NewCmd.AddCommand(commands.NewFuncCmd, commands.NewBotCmd)
	rootCmd.Execute()
}