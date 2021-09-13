package dispatcher

import (
	"crypto/tls"
	"net/http"
	"time"
)

var (
	Client = http.Client{
		Transport:     &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout:       time.Second * 30,
	}
)
