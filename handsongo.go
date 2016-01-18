package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
)

var (
	app       = cli.NewApp()
	header, _ = base64.StdEncoding.DecodeString("" +
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

func init() {
	fmt.Print(string(header))
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stderr)
}

func main() {

	app.Name = "handsongo"
	app.Version = "0.0.1"
	app.Usage = "handsongo service launcher"

	app.Commands = make([]cli.Command, 0)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Run error %q\n", err)
	}
}
