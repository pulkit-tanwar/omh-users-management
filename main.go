package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "omh-user-management-service"
	app.Usage = "omh-user-management-service"
	app.UsageText = "Create and manage users"

	app.Commands = []cli.Command{
		{
			Name:      "serve",
			Usage:     "Run an API Server to manage users",
			UsageText: "serve",
			Action:    RunServer,
		},
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

	err := app.Run(os.Args)
	if err != nil {
		log.Info("Error occurred:", err)
		os.Exit(1)
	}
}

func RunServer(*cli.Context) error {
	fmt.Println("TODO: code to run the HTTP Server")
	return nil
}
