package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"regexp"
)

const (
	DATA  = ".*?"
	SPACE = "\\s"
)

func main() {
	app := cli.NewApp()
	app.Author = "Nicolas Grange"
	app.Email = "grange74@gmail.com"
	app.Usage = "cli-log-parser filename"
	app.Version = "0.0.2"
	app.Action = parseLogFile
	app.Run(os.Args)
}

func parseLogFile(c *cli.Context) {
	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return
	}

	filename := c.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Example: 2015-03-18 08:54:07.498 [http-8080-6] ERROR Logger - executeParameterisedSql(): Error storing data in
	errorLogRegEx, err := regexp.Compile(DATA + SPACE + "\\[" + DATA + "\\]" + SPACE + "ERROR Logger" + DATA)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if errorLogRegEx.MatchString(line) {
			fmt.Println("ERROR: " + line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
