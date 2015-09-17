package main

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/tpyolang/tpyo-cli"
)

// main is the entrypoint
func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Tpyo atuhros"
	app.Email = "https://github.com/tpyolang/tpyo-cli"
	app.Version = "1.0.0"
	app.Usage = "Mkae tpyos"

	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	tpyo := tpyo.Tpyo{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(tpyo.Enocde(line))
	}
}
