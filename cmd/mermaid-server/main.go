package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/brunoluiz/mermaid-server/tmpl"
	"github.com/fsnotify/fsnotify"
	"github.com/jaschaephraim/lrserver"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mermaid-server",
		Usage: "⚡️ Serve Mermaid (MMD) files with live reloading",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Aliases: []string{"a"},
				Value:   ":7777",
				Usage:   "server hostname",
			},
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Required: true,
				Usage:    "mmd file path",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "mmd config json file path",
			},
		},
		Action: func(c *cli.Context) error {
			var err error
			input := c.String("input")
			address := c.String("address")
			configPath := c.String("config")

			// Load config json file
			config, err := loadConfig(configPath)
			if err != nil {
				return err
			}

			// Create file watcher
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				return err
			}
			defer watcher.Close()

			// Add dir to watcher
			if err = watcher.Add(input); err != nil {
				return err
			}

			// Create and start LiveReload server
			lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
			lr.SetStatusLog(nil)
			lr.SetErrorLog(nil)
			go lr.ListenAndServe() // nolint

			// Start goroutine that requests reload upon watcher event
			go func() {
				for {
					select {
					case event := <-watcher.Events:
						lr.Reload(event.Name)
					case err := <-watcher.Errors:
						logrus.Error(err)
					}
				}
			}()

			// Start serving html
			http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
				mmdDiagram, err := ioutil.ReadFile(input)
				if err != nil {
					logrus.Fatal(err)
				}

				t, err := template.New("mermaid").Parse(tmpl.IndexHTML)
				if err != nil {
					logrus.Error(err)
				}

				if err := t.Execute(rw, tmpl.IndexHTMLSubs{
					Content: string(mmdDiagram),
					Config:  template.JS(config),
				}); err != nil {
					logrus.Error(err)
				}
			})

			logrus.Infof("serving %s at %s", input, address)

			return http.ListenAndServe(address, nil)
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func loadConfig(configPath string) (string, error) {
	if configPath == "" {
		return "{}", nil
	}

	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	if !isJSON(string(config)) {
		return "", errors.New(configPath + " is not a valid json")
	}

	return string(config), nil
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
