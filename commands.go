package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/motemen/ghq/utils"
)

const BASE_URL = "https://api.chatwork.com/v1"

var Commands = []cli.Command{
	commandSend,
}

var commandSend cli.Command = cli.Command{
	Name:    "send",
	Aliases: []string{"g"},
	Usage:   "Send a message to chatwork",
	Action: func(c *cli.Context) error {
		room_id := c.Args().First()
		values := url.Values{}
		values.Add("body", "aaa")

		client := http.Client{}
		req, err := http.NewRequest(
			"POST",
			BASE_URL+"/rooms/"+room_id+"/messages",
			strings.NewReader(values.Encode()))
		utils.DieIf(err)

		req.Header.Add("X-ChatWorkToken", api_key())

		resp, err := client.Do(req)
		utils.DieIf(err)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		return nil
	},
}

func api_key() string {
	api_key := os.Getenv("CHATWORK_API_KEY")
	if api_key == "" {
		fmt.Println("please set CHATWORK_API_KEY")
		os.Exit(1)
	}

	return api_key
}
