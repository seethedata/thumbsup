//getApprovers.go gets approver audit trail for a given application
package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

//Application represents an application submission
type Approval struct {
	Name         string `json:name`
	Role         string `json:role`
	ApprovalDate string `json:approvalDate`
}

func getApprovers(addr common.Address) (appl []Approval, appName string) {
	conn, err := ethclient.Dial(blockchainAPI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contract, err := NewContract(addr, conn)
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}

	numApprovers, err := contract.GetNumApprovers(nil)
	if err != nil {
		log.Fatalf("Failed to get number of approvers: %v", err)
	}
	max := int(numApprovers.Int64())

	for i := 0; i < max; i++ {
		aName, aRole, aDate, err := contract.GetApprover(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("Failed to get approver at index: %v: %v", i, err)
		}
		appl = append(appl, Approval{Name: aName, Role: aRole, ApprovalDate: aDate})
	}
	appName, err = contract.GetApplicationName(nil)
	if err != nil {
		log.Fatalf("Failed to get application name: %v", err)
	}
	return appl, appName
}
