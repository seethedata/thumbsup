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
	"time"
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
	conn, err := ethclient.Dial(blockchainAPI)
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
	address, tx, _, err := DeployContract(auth, conn, common.HexToAddress(os.Getenv("USERADDRESS")), size, name)
	if err != nil {
		log.Fatalf("Failed to deploy new contract: %v", err)
	}
	log.Printf("Deploying contract at address %s in transaction 0x%x\n", address.String(), tx.Hash())
	time.Sleep(250 * time.Millisecond)

	contract, err := NewContract(address, conn)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}

	session := &ContractSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:   auth.From,
			Signer: auth.Signer,
			//GasLimit: big.NewInt(maxGas),
		},
	}
	log.Printf("MaxGas is %v\n", maxGas)
	// Add required approvers to contract. It returns tx and err
	for n, a := range r.Form["approverType"] {
		tx, err := session.AddRequiredRole(a)
		log.Printf("%s | %s | AddApp Tx hash is 0x%x\n", n, a, tx.Hash())
		if err != nil {
			log.Fatalf("Failed to add approver: %v", err)
		}
		time.Sleep(5000 * time.Millisecond)
	}

	metaContract, err := NewMetaContract(common.HexToAddress(os.Getenv("METACONTRACT")), conn)
	if err != nil {
		log.Fatalf("Failed to bind metacontract %v", err)
	}

	// Update metacontract. It returns tx ,err
	mtx, err := metaContract.AddApplication(auth, address, name)
	log.Printf("Meta Tx hash is 0x%x\n", mtx.Hash())
	if err != nil {
		log.Fatalf("Failed to add application to metacontract: %v", err)
	}

	http.Redirect(w, r, mainURL+"/submitted", http.StatusFound)
}
