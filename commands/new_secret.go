package commands

import (
	"fmt"
	"github.com/discless/discless/types/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	ffile string
	git  bool
)

var NewSecretCmd = &cobra.Command{
	Use: "secret [secret key] [secret]",
	Short: "Create a new bot with given name",
	Args: cobra.MinimumNArgs(2),
	RunE: FNewSecret,
}

func INewSecret() {
	NewSecretCmd.Flags().BoolVarP(&git, "git","g", false, "adds file to gitignore")
	NewSecretCmd.Flags().StringVarP(&ffile,"file","f","secret.yaml","defines the file name for the secret to be stored in")
}

func FNewSecret(command *cobra.Command, args []string) error {
	secret := &config.Secret{}

	if _, err := os.Stat(ffile); err == nil {
		f, err := ioutil.ReadFile(ffile)

		if err != nil {
			return err
		}

		err = yaml.Unmarshal(f,secret)

		if secret.Secrets == nil {
			secret.Secrets = make(map[string]string)
			secret.Secrets[args[0]] = args[1]
		}

		secret.Secrets[args[0]] = args[1]
	} else if os.IsNotExist(err) {
		secret.Secrets = make(map[string]string)
		secret.Secrets[args[0]] = args[1]
	} else {
		return err
	}

	mars, err := yaml.Marshal(secret)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ffile, mars, 0644)

	if err != nil {
		return err
	}

	fmt.Println("Created secret in", ffile)
	return nil
}

