//approveForm.go handles the form entery page
package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
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

func approveHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ethclient.Dial("http://" + blockchainAPI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(os.Getenv("USERKEY")), os.Getenv("USERPASSWORD"))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	r.ParseForm()
	contractAddress := r.Form["address"][0]
	aType, err := strconv.ParseInt(r.Form["approverType"][0], 10, 64)
	approverType := big.NewInt(aType)
	if err != nil {
		log.Fatalf("Failed to convert approverType: %v", err)
	}

	contract, err := NewContract(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}
	// Deploy contract
	_, err = contract.AddApprover(auth, common.HexToAddress(os.Getenv("USERADDRESS")), approverType)
	if err != nil {
		log.Fatalf("Failed to register approval: %v", err)
	}

	http.Redirect(w, r, mainURL+"/list", http.StatusFound)

}
