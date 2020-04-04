package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func createCliApp() *cli.App {
	bucketTreemapCommand := &cli.Command{
		Name: "treemap",
		Usage: "create count and size based treemaps for bucket prefixes",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "bucket",
				Usage: "bucket name",
			},
			&cli.IntFlag{
				Name: "depth",
				Usage: "maximum prefix depth",
				Value: 32,
			},
		},
	}

	return &cli.App{
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "bucket",
				Usage: "bucket commands",
				Subcommands: []*cli.Command{
					bucketTreemapCommand,
				},
			},
		},
		EnableBashCompletion: true,
	}
}

func main() {
	app := createCliApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
