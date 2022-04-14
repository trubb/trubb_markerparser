package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var fileName string

func main() {

	app := &cli.App{

		Name:  "Markerparser",
		Usage: "Translate between marker and artillery target/bookmark arrays",
		Commands: []*cli.Command{
			{
				Name:    "tuntematonFireSupport",
				Aliases: []string{"tfs"},
				Usage:   "from sweet marker array to tuntematon firesupport trps",
				Action:  sweetToTun,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "sourceFile",
						Aliases:     []string{"s"},
						Usage:       "load array from `FILE`",
						Destination: &fileName,
					},
				},
			},
			{
				Name:    "sweetMarkerSystem",
				Aliases: []string{"sms"},
				Usage:   "from tuntematon firesupport trps to sweet marker array",
				Action:  tunToSweet,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "sourceFile",
						Aliases:     []string{"s"},
						Usage:       "load array from `FILE`",
						Destination: &fileName,
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err)
	}
}

func tunToSweet(c *cli.Context) error {
	log.Printf("Trying to run tunToSweet with the filename %q", fileName)
	return nil
}

func sweetToTun(c *cli.Context) error {
	log.Printf("Trying to run sweetToTun with the filename %q", fileName)
	return nil
}

func writeToFile(fileName string, input string) error {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := file.WriteString(input + "\n"); err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func readFromFile(fileName string) error {

	byteArr, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	_ = byteArr
	// do something with, or return, byteArr

	return nil
}
