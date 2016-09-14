package main

import (
	"net/http"
	"io"
	"sync"
)

type (
	Conf struct {
		Http *http.Client
		WG *sync.WaitGroup
		Root string
		Verbose bool
	}

	Anchor struct {
		Url string
	}

	HtmlPage struct {
		Stream io.Reader
	}
)