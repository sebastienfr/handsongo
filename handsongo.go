package main

import (
	"encoding/base64"
	"fmt"
	logger "github.com/Sirupsen/logrus"
	"github.com/sebastienfr/handsongo/dao"
	"github.com/sebastienfr/handsongo/utils"
	"github.com/sebastienfr/handsongo/web"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"time"
)

var (
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string

	port               = 8020
	logLevel           = "warning"
	db                 = "mongodb://mongo/spirits"
	logFormat          = utils.TextFormatter
	statisticsDuration = 20 * time.Second

	header, _ = base64.StdEncoding.DecodeString(
		"DQouLCwsLC4uLi4gIC4gIC4uICAgTk1NTU0sLi4gLi4uLi4gICAgLi4uTU0sLiAuICAgICAgLi4sLCwsLi4NCjo3JCRPOi4gLi4uLi4uICAuIDdNSSAg" +
			"Li4gICAgIC4gICAgLiAgICArSSAuICAgICAgLi4uLk9aJCQuLg0KOiQkIC4gICAgKzhORDc6Ljo9OE1JPT06IDp+fn5+fn5+fiwuLj09PT09Li4uPTouOiR" +
			"aPy4gLiA6JC4gDQo6JCQgICAuPU04TU1NRC4uPU1OTk5NTT0gSSQkJCQkJCQkOiAsTU5NOE0uLiBNKz1ETU1EOiAuLjokLiANCjokJCAgLi5NTS4uLi4uIC" +
			"AgIDhNfi4gLi4uIC4gICAgIC4uLiAuICBNTS4uLk1NRCAuLiAuICAgOiQuIA0KOiQkIC4gIDdNRE1afiAgIC4gOE1+LiAgLjo6Ojo6Ojo6Oi4uIC4gIE1NL" +
			"i4gTSsgLiAgICAgICA6JC4uDQo6JCQgICAgLi4kTU04TTcgLiA4TX4uICAgPyQkJCQkJCQ3OiAuLiAgTU0uLiBNKyAgICAgICAgIDo3LiANCjokJCAgIC4u" +
			"IC4uLi5NOD0uIDhNfi4gICAuLi4gICAgIC4gLi4uICBNTS4uIE0rIC4gICAgICAgOiQuLg0KOiQkICAuLj06ICAuLE04Oi4gOE1+LiAgLiwsLCwsLCwsLC4" +
			"uLi4gIE1NLi4uTT8gLiAgIC4gLi46JC4gDQo6JCQgLiAuTThNTU1ETVouLiA4TX4uICAuPyQ3Nzc3Nzc3Oi4uLi4gTU0uLiBNKyAgICAgLi4uIDokLiANCj" +
			"okJC4gLi4uLi4uLi4uLiAgICAuICAgIC4uLi4gICAgICAuLi4uLi4gIC4gICAgLiAgICAuLi4uOiQuIA0KOlpaJE86Li4uICAgICAgICAgICAgICAgICAgI" +
			"CAgICAgICAgICAgICAgICAgICAgICAgIC4uT1paWi4uDQouLi4uLi4uLi4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgLi4u" +
			"LiAuIA0K")
)

func main() {
	// new app
	app := cli.NewApp()
	app.Name = "handsongo"
	app.Usage = "handsongo service launcher"

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	app.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	app.Authors = []cli.Author{cli.Author{Name: "sfr"}}
	app.Copyright = "Sfeir " + strconv.Itoa(time.Now().Year())

	// command line flags
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Value: port,
			Name:  "port",
			Usage: "Set the listening port of the webserver",
		},
		cli.StringFlag{
			Value: db,
			Name:  "db",
			Usage: "Set the mongo database connection string",
		},
		cli.StringFlag{
			Value: logLevel,
			Name:  "logl",
			Usage: "Set the output log level (debug, info, warning, error)",
		},
		cli.StringFlag{
			Value: logFormat,
			Name:  "logf",
			Usage: "Set the log formatter (logstash or text)",
		},
		cli.DurationFlag{
			Value: statisticsDuration,
			Name:  "statd",
			Usage: "Set the token duration (ex : 1h, 2h30m, 30s, 300ms)",
		},
	}

	// main action
	// sub action are also possible
	app.Action = func(c *cli.Context) error {
		// print header
		fmt.Println(string(header))

		// parse parameters
		port = c.Int("port")
		db = c.String("db")
		logLevel = c.String("logl")
		logFormat = c.String("logf")
		statisticsDuration = c.Duration("statd")

		fmt.Print("* --------------------------------------------------- *\n")
		fmt.Printf("|   port                    : %d\n", port)
		fmt.Printf("|   db                      : %s\n", db)
		fmt.Printf("|   logger level            : %s\n", logLevel)
		fmt.Printf("|   logger format           : %s\n", logFormat)
		fmt.Printf("|   statistic duration(s)   : %0.f\n", statisticsDuration.Seconds())
		fmt.Print("* --------------------------------------------------- *\n")

		// init log options from command line params
		err := utils.InitLog(logLevel, logFormat)
		if err != nil {
			logger.Warn("error setting log level, using debug as default")
		}

		webServer, err := web.BuildWebServer(db, dao.DAOMongo, statisticsDuration)

		if err != nil {
			return err
		}

		// serve
		webServer.Run(":" + strconv.Itoa(port))

		return nil
	}

	// run the app
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Run error %q\n", err)
	}
}
