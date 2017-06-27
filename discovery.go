//bcl.go is a demo of using blockchain for securities management
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type cfServices struct {
	Services []cfService `json:"services"`
}

type cfService struct {
	Platform string `json:"platform"`
	Name     string `json:"service"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func getServiceCreds() (string, string, string, string, string) {
	var cfServices cfServices
	fmt.Println("Starting")
	file, err := ioutil.ReadFile("./services.json")
	check("Read services JSON", err)

	err = json.Unmarshal(file, &cfServices)
	check("Unmarshal", err)

	env, _ := cfenv.Current()
	mainURL = "http://" + env.ApplicationURIs[0]
	blockchainAPI = os.Getenv("BLOCKCHAINAPI")
	services := env.Services

	var credentials map[string]interface{}
	var host string
	var password string
	var port string

	for _, service := range cfServices.Services {
		if _, ok := services[service.Name]; ok {
			credentials = services[service.Name][0].Credentials
			if _, ok := credentials[service.Host]; ok {
				host = credentials[service.Host].(string)
			} else {
				log.Fatal("Unable to identify Redis host from config. Platform attempted:" + service.Platform)
			}
			if _, ok := credentials[service.Password]; ok {
				password = credentials[service.Password].(string)
			} else {
				log.Fatal("Unable to identify Redis password from config. Platform attempted:" + service.Platform)
			}
			if _, ok := credentials[service.Port]; ok {
				switch credentials[service.Port].(type) {
				case string:
					port = credentials[service.Port].(string)
				case float64:
					port = strconv.FormatFloat(credentials[service.Port].(float64), 'f', -1, 64)
				default:
					log.Fatal("Redis port value is of unexpected type.")
				}
			} else {
				log.Fatal("Unable to identify Redis port from config. Platform attempted:" + service.Platform)
			}
			break
		}
	}
	return mainURL, host, password, port, blockchainAPI
}
