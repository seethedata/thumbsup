//audit.go serves the audit page
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

	t, err := template.ParseFiles("templates/audit.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.Approvals = getApprovers(common.HexToAddress(address))
	data.Applications = append(data.Applications, getApplication(common.HexToAddress(address)))
	t.Execute(w, data)
}
