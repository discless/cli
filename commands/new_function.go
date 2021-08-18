package commands

import (
	"fmt"
	"github.com/discless/discless-cli/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var NewFuncCmd = &cobra.Command{
	Use: "function [function name]",
	Short: "Create a new function with given name",
	Args: cobra.MinimumNArgs(1),
	RunE: FNewFunc,
}

func FNewFunc(cmd *cobra.Command, args []string) error {
	yamlTemplate := &types.Config{
		map[string]types.Function{
			args[0]:{
				File:     args[0]+".go",
				Function:	"Handler",
				Category: 	"",
			},
		},
	}
	mrs, _ := yaml.Marshal(yamlTemplate)
	ioutil.WriteFile("function.yaml",mrs,0644)
	goTemplate := `package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/discless/discless/types"
)

func ` + yamlTemplate.Functions[args[0]].Function + `(self *types.Self, s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
	return nil
}`
	ioutil.WriteFile(yamlTemplate.Functions[args[0]].File,[]byte(goTemplate),0644)

	fmt.Println("Created the function",args[0],"\nEdit its configuration in function.yaml or edit the function in",args[0]+".go")
	return nil
}
