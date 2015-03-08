package main

import (
	"C"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/bondj/slslack/config"
	"github.com/bondj/slslack/slack"
	"github.com/bondj/slslack/softlayer"
	"log"
	"net/http"
)

const sltoken = "DgnNlmIN6VpmiqhPZRjmiH9n"

var cfg = config.LoadConfig()

var (
	client *http.Client
	pool   *x509.CertPool
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		r.ParseForm()

		var cn string
		var cmd string
		var token string

		for key, values := range r.Form { // range over map
			for _, value := range values { // range over []string
				fmt.Println(key, value)
				if key == "channel_name" {
					cn = value
				} else if key == "command" {
					cmd = value
				} else if key == "token" {
					token = value
				}
			}
		}

		if token != cfg.Slack.Slacktoken {
			http.Error(w, "Invalid Request", 401)
		}

		if r.Method == "POST" {
			go func() {
				fmt.Println("Channel ", cn)
				fmt.Println("Command ", cmd)

				var m slack.Message
				m.Text = softlayer.GetNextInvoiceTotalAmount(client, cfg.Softlayer.User, cfg.Softlayer.Key)
				m.Channel = "#" + cn
				slack.SendMessage(client, m, cfg.Slack.Target)
			}()
		} else {
			http.Error(w, "Invalid request method.", 405)
		}
	})

	pool = x509.NewCertPool()
	fmt.Println(pool.AppendCertsFromPEM(pemCerts))

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		},
	}
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
