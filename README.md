# Discless CLI [![Go](https://github.com/discless/discless-cli/actions/workflows/go.yml/badge.svg)](https://github.com/discless/discless-cli/actions/workflows/go.yml)<br>
Discless CLI is a Command Line Interface to communicate with the Discless backend. It can be used to create discord bots and commands in a FaaS, so you don't have to worry about the command handling.

## Setup
First download the Discless binary from the ![releases](https://github.com/discless/discless-cli/releases),
or clone the repository using `git clone https://github.com/discless/discless-cli.git` and run `go build .` to compile Discless yourself
To start Discless, run
```shell
$ ./discless-cli start
Succesfully started Docker daemon.
```

## Examples
### Create a bot
First, run the `new bot` command and enter your bots token
```shell
$ ./discless-cli new bot <bot name> <prefix>
Created bot in bot.yaml
```
Your bot should be up and running now, time to create your first command.

## Run your bot
First, create a new secret for your token
```shell
$ ./discless-cli new secret token NDMyMTkx...
Created secret in secret.yaml
```

To use the token in your bots configuration, open your bots configuration and change the following
```
- token: 
+ token: secret.token
```

Now you can run your bot
```shell
$ ./discless-cli up bot.yaml
<bot-name> is running
```

Run
```shell
$ ./discless-cli new function <function name>
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
$ ./discless-cli deploy <bot name> <function configuration>.yaml
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
