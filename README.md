# Discless CLI
CLI to communicate with the Discless backend.

## Setup
To start Discless, run
```shell
$ discless-cli start
Succesfully started Docker daemon.
```

## Examples
### Create a bot and function
First, run the `new bot` command and enter your bots token
```shell
$ discless-cli new bot <bot name> <prefix>
Enter your bot's token: 
Succesfully created the bot <bot name>
```
Your bot should be up and running now, time to create your first command.

Run
```shell
$ discless-cli new function <function name>
Created the function <function name>
Edit its configuration in function.yaml or edit the function in <function name>.go
```
This creates a configuration file for the function (`function.yaml`) and a golang file that looks like
```go
package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/discless/discless/types"
)

func Handler(self *types.Self, s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
	return nil
}
```
You can freely edit this file.

To get your command up and running on the bot, run
```shell
$ discless-cli apply <bot name> <function configuration>.yaml
Succesfully uploaded the <function name> command
```

## Manually installing discless
To manually install discless, clone the repository
```shell
$ git clone https://github.com/discless/discless.git && cd discless && go run .
```
or use `go get`
```shell
go get https://github.com/discless/discless && cd $GOPATH/github.com/discless/discless && go run .
```