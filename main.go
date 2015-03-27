package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Author = "Nicolas Grange"
	app.Email = "grange74@gmail.com"
	app.Usage = "cli-log-parser filename"
	app.Version = "0.0.1"
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
