package main

import (
	"github.com/Sirupsen/logrus"
	"net/http"
)

func Run(config Conf){
	config.WG.Add(1)
	defer config.WG.Done()

	client := config.Http

	channel := make(chan HtmlPage)
	go RunParser(config, channel)

	resp, err := client.Get(config.Root)

	if err != nil {
		logrus.WithField("uri", config.Root).WithError(err).Error("Request failed to root url! Are you sure it's correct?")
		return
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("status", resp.StatusCode).WithField("uri", config.Root).Error("Request to root url was not OK!")
		return
	}

	channel <- HtmlPage{resp.Body}

	close(channel)
}