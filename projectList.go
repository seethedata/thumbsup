//projectList.go lists the status of the projects in the system
package main

import (
	"html/template"
	"net/http"
)

func projectListHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/projectList.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.BlockchainAPI = blockchainAPI
	t.Execute(w, data)
}
