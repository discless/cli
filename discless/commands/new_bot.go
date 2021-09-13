package commands

import (
	"fmt"
	"github.com/discless/discless/types/config"
	"github.com/discless/discless/types/kinds"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"syscall"
)

var (
	prompt		bool
	token		string
	name		string
	prefix		[]string
	file		string
)

var NewBotCmd = &cobra.Command{
	Use:   "bot [bot name]",
	Short: "Create a new bot with given name",
	Args:  cobra.MinimumNArgs(1),
	RunE:  FNewBot,
}

func INewBot()  {
	NewBotCmd.Flags().BoolVar(&prompt,"prompt",false, "opens a prompt to enter your bot's token")
	NewBotCmd.Flags().StringArrayVarP(&prefix, "prefix","p",nil,"set the prefix for the bot")
	NewBotCmd.Flags().StringVarP(&file, "file","f","bot.yaml","set the file name for the bot to be saved in")
}

func FNewBot(c *cobra.Command, args []string) error {
	bot := &config.Bot{
		Kind:     	kinds.Bot,
		Name:     	args[0],
		Prefix: 	"",
	}

	if prompt {
		fmt.Print("Enter your bot's token: ")
		btoken, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		bot.Token = string(btoken)
	}

	if prefix != nil {
		bot.Prefix = prefix[0]
	}

	enc, err := yaml.Marshal(bot)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, enc, 0644)

	if err != nil {
		return err
	}

	fmt.Println("✅ Created bot in", file)

	return nil
}