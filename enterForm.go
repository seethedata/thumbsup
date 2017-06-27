//enterForm.go handles the form entery page
package main

import (
	"html/template"
	"net/http"
)

func enterFormHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/enterForm.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.BlockchainAPI = blockchainAPI
	t.Execute(w, data)
}
