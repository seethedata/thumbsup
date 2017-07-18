//getApprovers.go gets approver audit trail for a given application
package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func applicationAuditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["applicationID"]

	t, err := template.ParseFiles("templates/approverList.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.Approvals, data.ApplicationName = getApprovers(common.HexToAddress(address))
	t.Execute(w, data)
}
