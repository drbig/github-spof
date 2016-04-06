package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type GitHubStatus struct {
	Status    string    `json:"status"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

func main() {
	args := os.Args[1:]
	var wg sync.WaitGroup
	wg.Add(1)

	var token string
	var channel string
	var ping int

	cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
	cmdFlags.StringVar(&token, "t", "", "Slack API Token")
	cmdFlags.StringVar(&channel, "c", "general", "Channel")
	cmdFlags.IntVar(&ping, "p", 10, "Message")

	if err := cmdFlags.Parse(args); err != nil {
		fmt.Printf("All is wrong")
		os.Exit(1)
	}
	channel, _ = url.QueryUnescape(channel)
	ticker := time.NewTicker(time.Duration(rand.Intn(ping)) * time.Second)
	quit := make(chan struct{})
	go func(token string, channel string) {
		currentMessage := ""
		for {
			select {
			case <-ticker.C:
				githubResponse, err := http.Get(fmt.Sprintf("https://status.github.com/api/last-message.json"))
				if err != nil {
					fmt.Printf("%s", err)
					continue
				}
				var m GitHubStatus
				byt, err := ioutil.ReadAll(githubResponse.Body)
				if err != nil {
					fmt.Printf("%s", err)
					continue
				}
				err = json.Unmarshal(byt, &m)
				if err != nil {
					fmt.Printf("%s", err)
					continue
				}
				text, _ := url.QueryUnescape(fmt.Sprintf("[%s] %s. by GitHub Status", m.Status, m.Body))
				if text != currentMessage {
					_, err := http.Get(fmt.Sprintf("https://slack.com/api/chat.postMessage?token=%s&text=%s&channel=%s", token, text, channel))
					if err != nil {
						fmt.Printf("%s", err)
					}
					currentMessage = text
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}(token, channel)
	wg.Wait()
}
