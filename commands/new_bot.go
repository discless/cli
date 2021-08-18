package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/discless/discless/types"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"net/http"
	"syscall"
)

var NewBotCmd = &cobra.Command{
	Use: "bot [bot name] [prefix]",
	Short: "Create a new bot with given name",
	Args: cobra.MinimumNArgs(2),
	RunE: FNewBot,
}

func FNewBot(c *cobra.Command, args []string) error {
	bot := &types.Self{
		Name: args[0],
		Prefix:args[1],
	}
	bdy, err := json.Marshal(bot)

	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:6969/bot", bytes.NewBuffer(bdy))
	if err != nil {
		return err
	}

	fmt.Print("Enter your bot's token: ")
	btoken, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}


	token := string(btoken)

	req.Header.Add("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Succesfully created the bot", args[0])
	return nil
}