package main

import (
	"fmt"
  "os"
  "time"
  "bytes"
  "net/http"
  "encoding/json"

  "github.com/urfave/cli"
)

var webhook string

func toSlack(channelName string, message string) {
  type payload struct {
    Channel string `json:"channel"`
    Text string `json:"text"`
  }

  jsonStr, err := json.Marshal(payload{ Channel: channelName, Text: message })
  req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
  if err != nil {
      panic(err)
  }
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close() 

  timestamp := time.Now().Format(time.RFC1123)
  fmt.Printf("%s: Message sent to channel %s\n", timestamp, channelName)
}

func main() {
	app := cli.NewApp()
  app.Name = "slack-notifier"
  app.Usage = "slack-notifier -c <channel_name> -m <message>"
  app.Version = "0.0.1"

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "channel, c",
      Usage: "channel where to send the message",
    },
    cli.StringFlag{
      Name: "message, m",
      Usage: "message to send",
    },
  }

  app.Action = func(c *cli.Context) error {
    channelName := c.String("channel")
    message := c.String("message")

    if !c.IsSet("channel") || !c.IsSet("message") {
      fmt.Println("Usage:", app.Usage)
      fmt.Println("\nRun \"slack-notifier help\" for more info")
      return nil
    }

    toSlack(channelName, message)

    return nil
  }

  app.Run(os.Args)
}
