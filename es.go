package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
	_ "fmt"
	"github.com/elastic/go-elasticsearch/v9"
)

var EsClient *elasticsearch.Client
func InitEs() error{
	cfg := elasticsearch.Config{
	Transport: &http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second,
		DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
			},
		},
	}
	var err error
	EsClient, err = elasticsearch.NewClient(cfg)
	if err != nil{
		return err
	}
	res, err := EsClient.Info()
	if err != nil{
		return err
	}
	defer res.Body.Close()
	// fmt.Println(res)
	return nil
}