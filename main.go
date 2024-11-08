package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
)



func main() {
	app := &cli.App{
		Name:  "HealthChecker",
		Usage: "My CLI app",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "domain",
				Aliases: []string{"d"},
				Usage:   "Domain name to check",
				Required: false,
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Port to check on",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" { // fixed missing braces for the block
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err) // fixed indentation
	}
}
