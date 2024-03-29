package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/pulkit-tanwar/omh-users-management/lib/api"
	"github.com/pulkit-tanwar/omh-users-management/lib/config"
	"github.com/pulkit-tanwar/omh-users-management/lib/database"
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
			UsageText: "serve -e ENV -p PORT --host HOST --api-path API_PATH",
			Action:    RunServer,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:   "env, e",
					Value:  config.DefaultEnv,
					Usage:  "environment (dev | test | stage | load | prod)",
					EnvVar: "ENV",
				},
				&cli.StringFlag{
					Name:   "host",
					Value:  config.DefaultHost,
					Usage:  "The host to listen on.",
					EnvVar: "HOST",
				},
				&cli.StringFlag{
					Name:   "api-path",
					Value:  config.DefaultAPIPath,
					Usage:  "url path prefix for mounting api router",
					EnvVar: "API_PATH",
				},
				&cli.IntFlag{
					Name:   "port,p",
					Value:  config.DefaultPort,
					Usage:  "port to listen on.",
					EnvVar: "PORT",
				},
			},
		},
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// log.SetReportCaller(true)

	err := app.Run(os.Args)
	if err != nil {
		log.Info("Error occurred:", err)
		os.Exit(1)
	}
}

func RunServer(c *cli.Context) error {
	cfg := config.NewConfig(c.String("env"), c.String("host"), c.Int("port"), c.String("api-path"))
	log.Infof("ENV:      %s", cfg.Env)
	log.Infof("HOST: %s", cfg.Host)
	log.Infof("PORT: %d", cfg.Port)
	log.Infof("API PATH: %s", cfg.APIPath)

	err := initDatabaseClient()
	if err != nil {
		return errors.Wrap(err, "Failed to start session")
	}

	if err := RunHTTPServer(cfg); err != nil {
		return err
	}

	return nil
}

// RunHTTPServer - Function that will run HTTP Server
func RunHTTPServer(cfg *config.Config) error {
	server := api.NewServer(cfg)
	err := server.Start()
	if err != nil {
		log.Errorf("Error while starting HTTP Server. Err:%+v", err)
		return errors.Wrap(err, "server.Start failed")
	}
	return nil
}

func initDatabaseClient() error {
	database.DB = &database.SQLDbClient{}
	return database.DB.DBConnect()
}
