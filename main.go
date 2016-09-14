package main

import (
	"github.com/urfave/cli"
	"os"
	"net/http"
	"sync"
	"errors"
	"github.com/fatih/color"
)

func main() {
	app := cli.NewApp()
	app.Name = "slinky"
	app.Email = "peter.clark@ft.com"
	app.Version = "0.0.1"
	app.Usage = "Checks links for abnormal responses in an html page."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "url",
			Usage: "Root `URL` to scan for links and check responses",
		},
		cli.BoolFlag{
			Name: "details, d",
			Usage: "Shows all responses (including OK ones)",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("url") == "" {
			return errors.New("Please specify a url using --url.")
		}

		color.Yellow("  Running Slinky!")
		color.Yellow("  ---------------")

		wg := &sync.WaitGroup{}
		Run(Conf{
			Http: &http.Client{},
			WG: wg,
			Root: c.String("url"),
			Verbose: c.Bool("details"),
		})
		wg.Wait()

		color.Yellow("\n  All done! If this is all you see, then everything is a-ok. If you don't believe me, try running with --details for more info.")
		return nil
	}

	app.Run(os.Args)
}
