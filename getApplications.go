//applicationList.go lists the status of the applications in the system
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

//Application represents an application submission
type Application struct {
	Name              string `json:name`
	Address           string `json:address`
	Status            string `json:status`
	Specification     string `json:specification`
	RequiredApprovals string `json:requiredApprovers`
	CurrentApprovals  string `json:currentApprovers`
}

func getApplications() (apps []Application) {
	conn, err := ethclient.Dial("http://" + blockchainAPI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	metaContract, err := NewMetaContract(common.HexToAddress(os.Getenv("METACONTRACT")), conn)
	if err != nil {
		log.Fatalf("Failed to convert address: %v", err)
	}

	numberOfApplications, err := metaContract.GetNumberApplications(nil)
	if err != nil {
		log.Fatalf("Failed to get number of applications: %v", err)
	}
	max := int(numberOfApplications.Int64()) + 1
	for i := 1; i < max; i++ {
		address, name, err := metaContract.GetApplication(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("Failed to get application %i: %v", i, err)
		}

		contract, err := NewContract(address, conn)
		if err != nil {
			log.Fatalf("Failed to get contract: %v", i, err)
		}

		var status string
		s, err := contract.GetStatus(nil)
		if err != nil {
			log.Fatalf("Failed to get status: %v", i, err)
		}
		switch int(s.Int64()) {
		case 1:
			status = "Submitted"
		case 2:
			status = "Valid"
		case 3:
			status = "Running"
		case 4:
			status = "Decommissioned"
		case 5:
			status = "Deployed"
		case 6:
			status = "Approved"
		case 7:
			status = "Rejected"
		default:
			status = "Unknown"
		}
		if err != nil {
			log.Fatalf("Failed to get status: %v", i, err)
		}

		cpu, ram, capacity, nic, err := contract.GetSpecification(nil)
		if err != nil {
			log.Fatalf("Failed to get specification: %v", i, err)
		}
		spec := fmt.Sprintf("CPU: %sGHz, RAM: %sGB, Capacity: %sGB, NICs: %s", cpu, ram, capacity, nic)
		required := ""
		current := ""
		_, req, cur, err := contract.GetApprovalsStatus(nil)
		for i := 0; i < 5; i++ {
			if int(req[i].Int64()) == 1 {
				if required == "" {
					required = approvers[i]
				} else {
					required = required + ", " + approvers[i]
				}
			}
			if int(cur[i].Int64()) == 1 {
				if current == "" {
					current = approvers[i]
				} else {
					current = current + ", " + approvers[i]
				}
			}
		}
		apps = append(apps, Application{Name: name, Address: address.String(), Status: status, Specification: spec, RequiredApprovals: required, CurrentApprovals: current})
	}
	return apps
}
