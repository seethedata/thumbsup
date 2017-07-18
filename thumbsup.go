//thumbsup.go is a blockchain demo for approvals
package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	pool          *redis.Pool
	mainURL       string
	blockchainAPI string
	metacontract  string
	router        *mux.Router
	approvers     [5]string
	maxGas        int
)

//Payload is a generic Container to hold data for templates
type Payload struct {
	URL             string        `json:"URL" redis:"URL"`
	Applications    []Application `json:applications redis:applications`
	Approvals       []Approval    `json:approvals redis:approvals`
	ApplicationName string        `json:applicationName redis:applicationName`
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
	approvers = [5]string{"Clinical", "Doctor", "Developer", "Application", "IT"}
	maxGas = 4700000
}

func main() {
	initialize()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/enter", enterFormHandler)
	router.HandleFunc("/create", createAppHandler)
	router.HandleFunc("/submitted", submittedHandler)
	router.HandleFunc("/", approveFormHandler)
	router.HandleFunc("/approve", approveHandler)
	router.HandleFunc("/approved", approvedHandler)
	router.HandleFunc("/list", applicationListHandler)
	router.HandleFunc("/{applicationID}/deploy", deployHandler)
	router.HandleFunc("/{applicationID}/audit", applicationAuditHandler)

	http.Handle("/images/", http.FileServer(http.Dir("/app")))
	http.Handle("/css/", http.FileServer(http.Dir("/app")))
	http.Handle("/fonts/", http.FileServer(http.Dir("/app")))
	http.Handle("/js/", http.FileServer(http.Dir("/app")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
