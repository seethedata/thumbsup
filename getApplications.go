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
	log.Printf("Found %s applications\n", numberOfApplications)
	max := int(numberOfApplications.Int64()) + 1
	for i := 1; i < max; i++ {
		address, name, err := metaContract.GetApplication(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("Failed to get application %i: %v", i, err)
		}

		contract, err := NewContract(address, conn)
		if err != nil {
			log.Fatalf("Failed to get contract: %v", err)
		}

		var status string
		s, numReq, numAppr, err := contract.GetApprovalsStatus(nil)
		if err != nil {
			log.Fatalf("Failed to get status: %v", err)
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
			log.Fatalf("Failed to parse status: %v", err)
		}

		cpu, ram, capacity, nic, err := contract.GetSpecification(nil)
		if err != nil {
			log.Fatalf("Failed to get specification: %v", err)
		}
		spec := fmt.Sprintf("CPU: %sGHz, RAM: %sGB, Capacity: %sGB, NICs: %s", cpu, ram, capacity, nic)
		required := ""
		current := ""
		for i := 0; i < int(numReq.Int64()); i++ {
			role, err := contract.GetRequiredRole(nil, big.NewInt(int64(i)))
			if err != nil {
				log.Fatalf("Failed to get required role %s: %v", i, err)
			}
			if i > 0 {
				required = required + ", "
			}
			required = required + role
		}

		log.Printf("numAppr is %s\n", numAppr.Int64())
		for i := 0; i < int(numAppr.Int64()); i++ {
			role, err := contract.GetApprovedRole(nil, big.NewInt(int64(i)))
			if err != nil {
				log.Fatalf("Failed to get approved role %s: %v", i, err)
			}
			if i > 0 {
				current = current + ", "
			}
			current = current + role
		}

		apps = append(apps, Application{Name: name, Address: address.String(), Status: status, Specification: spec, RequiredApprovals: required, CurrentApprovals: current})
	}
	return apps
}
