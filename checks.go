package main

import (
	"github.com/Sirupsen/logrus"
	"net/http"
	"github.com/fatih/color"
	"net/url"
	"strconv"
	"strings"
	"errors"
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

	url, err := buildUrl(config, anchor)
	if err != nil {
		return
	}

	resp, err := client.Get(url)

	if err != nil {
		logrus.WithField("url", anchor.Url).WithError(err).Error("Failed to make request!")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if config.Verbose {
			color.Green("  %-10s %10s  %s", "OK!", "200", anchor.Url)
		}
	} else {
		color.Red("  %-10s %10s  %s", "Not OK!", strconv.Itoa(resp.StatusCode), anchor.Url)
	}
}

func buildUrl(config Conf, anchor Anchor) (string, error) {
	uri, _ := url.Parse(config.Root)

	if strings.HasPrefix(anchor.Url, "#") {
		color.Yellow("  %-10s %10s  %-100s %-50s", "Skipped!", "", anchor.Url, "It looks like an in-page anchor.")
		return "", errors.New("Skipping because of #")
	}

	if strings.HasPrefix(anchor.Url, "http"){
		return anchor.Url, nil
	}

	uri.Path = anchor.Url
	return uri.String(), nil
}