//bcl.go is a demo of using blockchain for securities management
package main

import (
	//	"encoding/json"
	//	"fmt"
	//	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	//	"github.com/satori/go.uuid"
	//	"io/ioutil"
	"log"
	"net/http"
	//	"strconv"
	//	"strings"
	"time"
)

var (
	pool          *redis.Pool
	mainURL       string
	blockchainAPI string
	router        *mux.Router
)

//Product is a Dell product
type Product struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Image string `json:"image"`
	Price int    `json:"price"`
}

//Payload is a generic Container to hold data for templates
type Payload struct {
	URL           string `json:"URL" redis:"URL"`
	BlockchainAPI string `json:"blockchainAPI" redis:"blockchainAPI"`
	Products      []Product
}

func newPool(addr string, port string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr+":"+port)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func initialize() {
	var host string
	var password string
	var port string

	mainURL, host, password, port, blockchainAPI = getServiceCreds()

	pool = newPool(host, port, password)
}

func main() {
	initialize()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/projectList", projectListHandler)
	router.HandleFunc("/enter", enterFormHandler)
	router.HandleFunc("/", approveFormHandler)

	http.Handle("/images/", http.FileServer(http.Dir("/app")))
	http.Handle("/css/", http.FileServer(http.Dir("/app")))
	http.Handle("/fonts/", http.FileServer(http.Dir("/app")))
	http.Handle("/js/", http.FileServer(http.Dir("/app")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
