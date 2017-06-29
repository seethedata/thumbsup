//applicationList.go lists the status of the applications in the system
package main

import (
	"html/template"
	"net/http"
)

func applicationListHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/applicationList.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.BlockchainAPI = blockchainAPI
	t.Execute(w, data)
}
