//deploy.go shows the deployment of an approved application
package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
)

func deployHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["applicationID"]
	conn, err := ethclient.Dial(blockchainAPI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(os.Getenv("USERKEY")), os.Getenv("USERPASSWORD"))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contract, err := NewContract(common.HexToAddress(address), conn)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}
	// Register Approval
	_, err = contract.ChangeStatus(auth, big.NewInt(int64(5)))
	if err != nil {
		log.Fatalf("Failed to update contract status: %v", err)
	}

	t, err := template.ParseFiles("templates/deploy.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	t.Execute(w, data)
}
