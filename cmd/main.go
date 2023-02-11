package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var fileName string

// a trp represents one entry in the target reference point array
// that makes up the second subsection of the Tuntematon Fire Support
// bookmark + saved trp data structure
type trp struct {
	name   string    // target reference point name
	xCoord string    // because apparently there's string coordinates
	yCoord string    // because apparently there's string coordinates
	coords []float64 // actual coordinates
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

	// setup
	var trps []trp
	var trpString strings.Builder
	trpString.WriteString("[[],[")

	// read sweet marker array
	swmArr, err := readFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// clean it
	swmArr = strings.ReplaceAll(swmArr, "\n", "")
	swmArr = strings.ReplaceAll(swmArr, "\t", "")
	swmArr = strings.ReplaceAll(swmArr, " ", "")

	// break up into individual swmEntries
	swmEntries := strings.Split(swmArr, "],[")

	// extract the three first parts of each entry
	// and use them to create a valid tunEntry
	for _, entry := range swmEntries {
		var tunEntry trp

		// strip containing brackets
		entry := strings.ReplaceAll(entry, "[", "")
		entry = strings.ReplaceAll(entry, "]", "")

		// divide into string slice
		fields := strings.Split(entry, ",")

		// ignore any "" entry, which are generally a sweet marker's line
		// if it aint a line, whoever created the TRP is an idiot anyway
		if fields[0] == "\"\"" {
			continue
		}

		// set TRP name
		tunEntry.name = fields[0]

		// set detailed coordinates
		x, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return err
		}
		tunEntry.coords = append(tunEntry.coords, float64(x), float64(y))

		// wrangle the xCoord and yCoord string fields so that there's zeroes in front!
		for j := 1; j <= 2; j++ {
			if strings.Contains(fields[j], ".") {

				// discard decimals
				k := strings.Index(fields[j], ".")

				// add enough zeroes to the front of the string to pad up to 5 characters
				var sb strings.Builder
				pad := strings.Repeat("0", 5-len(fields[j][:k]))
				sb.WriteString("\"" + pad + fields[j][:k] + "\"")

				fields[j] = sb.String()
			}
		}
		tunEntry.xCoord = fields[1]
		tunEntry.yCoord = fields[2]

		trps = append(trps, tunEntry)
	}

	// write each entry into the output
	for _, trp := range trps {
		start := "["
		name := trp.name + ","
		xcoord := trp.xCoord + ","
		ycoord := trp.yCoord + ","
		coordArr := "[" + fmt.Sprintf("%g", trp.coords[0]) + "," + fmt.Sprintf("%g", trp.coords[1]) + "]"
		end := "],"
		trpString.WriteString(start + name + xcoord + ycoord + coordArr + end)
	}

	// remove dangling comma and close the array structure
	final := strings.TrimRight(trpString.String(), ",")
	final += "]]"

	log.Println("Parsed the input to the following trp array:")
	fmt.Println("\n" + final + "\n")

	// write to file
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

	byteArr, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return string(byteArr), nil
}
