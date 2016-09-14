package main

import (
	"github.com/Sirupsen/logrus"
	"net/http"
	"github.com/fatih/color"
	"net/url"
	"strconv"
)

func RunChecker(config Conf, channel chan Anchor){
	config.WG.Add(1)
	defer config.WG.Done()

	for {
		anchor, more := <- channel

		if !more {
			break
		}

		CheckAnchor(config, anchor)
	}
}

func CheckAnchor(config Conf, anchor Anchor){
	client := config.Http
	resp, err := client.Get(buildUrl(config, anchor))

	if err != nil {
		logrus.WithField("url", anchor.Url).WithError(err).Error("Failed to make request!")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if config.Verbose {
			color.Green("  OK! %10s:  %s", "200", anchor.Url)
		}
	} else {
		color.Red("  Not OK! %10s:  %s", strconv.Itoa(resp.StatusCode), anchor.Url)
	}
}

func buildUrl(config Conf, anchor Anchor) string {
	uri, _ := url.Parse(config.Root)
	uri.Path = anchor.Url
	return uri.String()
}