package commands

import (
	"bytes"
	"fmt"
	config2 "github.com/discless/discless/types/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var DeployCmd = &cobra.Command{
	Use: "deploy [bot] [function configuration]",
	Short: "Deploys a given function to a bot",
	Args: cobra.MinimumNArgs(2),
	RunE: FDeploy,
}

func FDeploy(c *cobra.Command, args []string) error {
	file, err := ioutil.ReadFile(args[1])
	if err != nil {
		return err
	}
	config := &config2.Config{}
	yaml.Unmarshal(file, config)

	for function, ap := range config.Functions {
		PostDeploy(function,ap, args[0])
	}

	return nil
}

func PostDeploy(name string, function config2.Function, bot string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
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

	req, err := http.NewRequest("POST", "http://localhost:8080/function", bytes.NewReader(body.Bytes()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}

	fmt.Println("Succesfully uploaded the", name, "command")

	return nil
}

