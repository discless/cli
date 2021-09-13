package commands

import (
	"bytes"
	"fmt"
	"github.com/discless/cli/discless"
	"github.com/discless/cli/discless/dispatcher"
	config2 "github.com/discless/discless/types/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy [bot] [function configuration]",
	Short: "Deploys a given function to a bot",
	Args:  cobra.MinimumNArgs(2),
	RunE:  FDeploy,
}

func FDeploy(c *cobra.Command, args []string) error {
	file, err := ioutil.ReadFile(args[1])
	if err != nil {
		return err
	}
	config := &config2.Config{}
	yaml.Unmarshal(file, config)

	for function, ap := range config.Functions {
		err := PostDeploy(function,ap, args[0])
		if err != nil {
			return err
		}
	}

	return nil
}

func PostDeploy(name string, function config2.Function, bot string) error {
	token, err := discless.GetKey(bot)
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("Name",name)
	writer.WriteField("Function", function.Function)
	writer.WriteField("Category", function.Category)
	writer.WriteField("Bot",bot)

	fw, err := writer.CreateFormFile("Function",function.File)
	if err != nil {
		return err
	}
	file, err := os.Open(function.File)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	writer.Close()

	req, err := http.NewRequest("POST", "https://" + discless.Host + ":" + discless.Port + "/function", bytes.NewReader(body.Bytes()))
	req.Header.Set("Authorization", "Basic " + token)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := dispatcher.Client.Do(req)

	b, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("❌ Request failed with response code: %d \n ⤷ %s", rsp.StatusCode, string(b))
		return nil
	}

	fmt.Println("✅ Successfully uploaded the", name, "command")

	return nil
}

