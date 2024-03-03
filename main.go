package main

import (
	"flag"
	"fmt"

	"github.com/NY-Daystar/populate-redis/configuration"
	"github.com/NY-Daystar/populate-redis/database"
	"github.com/NY-Daystar/populate-redis/utils"
	"go.uber.org/zap"
)

var cfg configuration.Config

var nbr *int
var flush *bool
var verbose *bool
var logger *zap.Logger

// init function to setup generation
func init() {
	nbr = flag.Int("n", 10, "Number of users to generate")
	flush = flag.Bool("f", false, "Whether to flush the database or not before generation")
	verbose = flag.Bool("v", false, "Make the process verbose or not")
}

// main Entrypoint of the program
func main() {
	flag.Parse()

	logger = utils.SetLogger(verbose)
	logger.Sugar().Infof("Project: %s", utils.Yellow(configuration.PROJECT_NAME))
	logger.Sugar().Infof("Version: %s", utils.Yellow(configuration.PROJECT_VERSION))

	configuration.ReadConfig(&cfg, logger)

	client := database.NewRedisClient(cfg.Redis)
	logger.Sugar().Debugf("Redis client created : %v", client)

	if *flush {
		fmt.Printf("%s", utils.Red("\t\tAre you SURE to flush this database ? "))
		if utils.AskForConfirmation() {
			logger.Sugar().Debug("Proceeding with the flush")
			database.FlushDatabase(client, logger)
			logger.Sugar().Debug(utils.Red("Database flushed"))
		} else {
			logger.Sugar().Debug(utils.Blue("No flush"))
		}
	}

	logger.Sugar().Infof("Generate %d users", *nbr)
	for i := 0; i < *nbr; i++ {
		database.AddUser(client, logger, *verbose)
	}
}
