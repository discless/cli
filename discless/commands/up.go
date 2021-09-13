package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/discless/cli/discless"
	"github.com/discless/cli/discless/dispatcher"
	"github.com/discless/cli/discless/util"
	"github.com/discless/discless/types/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

var (
	secretf []string
)

var UpCmd = &cobra.Command{
	Use:   "up [bot configuration file]",
	Short: "Run your bot",
	Args:  cobra.MinimumNArgs(1),
	RunE:  FUp,
}

func IUp() {
	UpCmd.Flags().StringSliceVarP(&secretf,"secrets","s",nil,"provide secret files to read")
}

func FUp(c *cobra.Command, args []string) error {
	secrets := make(map[string]string)

	if secretf != nil {
		for _,fn := range secretf {
			s := &config.Secret{
			}
			f, err := ioutil.ReadFile(fn)
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(f,s)
			if err != nil {
				return err
			}

			for key,val := range s.Secrets {
				secrets[key] = val
			}
		}
	}

	bot := &config.Bot{}

	f, err := ioutil.ReadFile(args[0])

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f,bot)

	if err != nil {
		return err
	}

	for _,el := range []*string{&bot.Token,&bot.Services.Database.MongoDB.Password,&bot.Services.Database.SQL.Password} {
		res, err := util.ReplaceSecret(*el,secrets)
		if err != nil {
			return err
		}
		*el = res
	}

	botb, err := json.Marshal(bot)

	req, err := http.NewRequest("POST", "https://" + discless.Host + ":" + discless.Port +"/bot", bytes.NewBuffer(botb))
	if err != nil {
		return err
	}

	resp, err := dispatcher.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("✅ Bot \"" + bot.Name + "\" is running")

	bodyb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	token := string(bodyb)

	if err := discless.AddKey(bot.Name, token); err != nil {
		return err
	}

	return nil
}
