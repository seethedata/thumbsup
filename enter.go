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

func enterFormHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/enterForm.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	t.Execute(w, data)
}

func submittedHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/submitted.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	t.Execute(w, data)
}

func createAppHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ethclient.Dial("http://" + blockchainAPI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(os.Getenv("USERKEY")), os.Getenv("USERPASSWORD"))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	r.ParseForm()
	s, err := strconv.ParseInt(r.Form["environmentSize"][0], 10, 64)
	if err != nil {
		log.Fatalf("Failed to convert size: %v", err)
	}
	size := big.NewInt(s)

	name := r.Form["applicationName"][0]

	// Deploy contract. It returns address, tx, contract, err
	address, _, _, err := DeployContract(auth, conn, common.HexToAddress(os.Getenv("USERADDRESS")), size, name)
	if err != nil {
		log.Fatalf("Failed to deploy new contract: %v", err)
	}
	log.Printf("Deploying contract at address %s\n", address.String())

	contract, err := NewContract(address, conn)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}

	// Add required approvers to contract. It returns tx and err
	for _, a := range r.Form["approverType"] {
		tx, err := contract.AddRequiredRole(auth, a)
		log.Printf("Tx is %s\n", tx)
		if err != nil {
			log.Fatalf("Failed to add approver: %v", err)
		}
	}

	metaContract, err := NewMetaContract(common.HexToAddress(os.Getenv("METACONTRACT")), conn)
	if err != nil {
		log.Fatalf("Failed to bind metacontract %v", err)
	}

	// Update metacontract. It returns tx ,err
	tx, err := metaContract.AddApplication(auth, address, name)
	log.Printf("Tx is %s\n", tx)
	if err != nil {
		log.Fatalf("Failed to add to metacontract: %v", err)
	}

	http.Redirect(w, r, mainURL+"/submitted", http.StatusFound)
}
