//approveForm.go handles the form entery page
package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func approveFormHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/approveForm.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.Applications = getApplications()
	t.Execute(w, data)
}

func approvedHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/approved.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	t.Execute(w, data)
}

func approveHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ethclient.Dial(blockchainAPI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(os.Getenv("USERKEY")), os.Getenv("USERPASSWORD"))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	r.ParseForm()
	contractAddress := r.Form["address"][0]
	approverType := r.Form["approverType"][0]
	approverName := r.Form["approverName"][0]

	contract, err := NewContract(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}
	// Register Approval
	_, err = contract.Approve(auth, approverName, approverType)
	if err != nil {
		log.Fatalf("Failed to register approval: %v", err)
	}

	http.Redirect(w, r, mainURL+"/approved", http.StatusFound)

}
