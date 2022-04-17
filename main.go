package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var fileName string

type tun struct {
	name   string
	xCoord string
	yCoord string
}

func main() {
	app := &cli.App{
		Name:  "markerparser",
		Usage: "Translate between marker and artillery trp arrays",
		Commands: []*cli.Command{
			{
				Name:    "toTuntematonFireSupport",
				Aliases: []string{"totfs"},
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// sweetToTun transforms a provided Sweet Markers array into Tuntematon Firesupport's target reference point format.
func sweetToTun(c *cli.Context) error {

	log.Printf("Running sweetToTun on file %q", fileName)

	var trps []tun
	var trpString strings.Builder
	trpString.WriteString("[[],[")

	content, err := readFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// clean
	content = strings.ReplaceAll(content, "\n", "")
	content = strings.ReplaceAll(content, "\t", "")
	content = strings.ReplaceAll(content, " ", "")

	// break up into individual entries
	entries := strings.Split(content, "],[")

	// extract the three first parts of each entry
	for _, str := range entries {
		entry := strings.ReplaceAll(str, "[", "")
		entry = strings.ReplaceAll(entry, "]", "")

		components := strings.Split(entry, ",")

		// ignore any "" entry, as those are impossible to reference accurately
		if components[0] == "\"\"" {
			continue
		}

		// wrangle the coordinates so that there's zeroes in front!
		for j := 1; j <= 2; j++ {
			if strings.Contains(components[j], ".") {

				// discard decimals
				k := strings.Index(components[j], ".")

				// add enough zeroes to the front of the string to pad up to 5 characters
				var sb strings.Builder
				pad := strings.Repeat("0", 5-len(components[j][:k]))
				sb.WriteString("\"" + pad + components[j][:k] + "\"")

				components[j] = sb.String()
			}
		}

		trps = append(trps, tun{components[0], components[1], components[2]})
	}

	for _, trp := range trps {
		trpString.WriteString("[" + trp.name + "," + trp.xCoord + "," + trp.yCoord + "],")
	}

	final := strings.TrimRight(trpString.String(), ",")
	final += "]]"

	log.Println("Parsed the input to the following trp array:")
	fmt.Println("\n" + final + "\n")

	outFile := createFile()
	err = writeToFile(outFile, final)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finished")

	return nil
}

// createFile creates a text file with a supposedly unique filename.
// If another file with the same name already exists, that file will be deleted.
func createFile() string {

	fileName := "parsed_markers_" + time.Now().Format("2006-01-02T15:04:05Z07:00") + ".txt" // RFC3339 format

	err := os.Remove(fileName)
	if err != nil {
		log.Println("No file with same name found, creating...")
	} else {
		log.Println("Deleted pre-existing file with same name")
	}

	_, err = os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created output file %v\n", fileName)

	return fileName
}

// writeToFile writes the provided string into the specified file.
func writeToFile(target string, input string) error {

	file, err := os.OpenFile(target, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

// readFromFile reads the content of a specified file and returns its contents as a string.
func readFromFile(fileName string) (string, error) {

	log.Printf("Reading from file %v\n", fileName)

	byteArr, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(byteArr), nil
}
