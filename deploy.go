//deploy.go shows the deployment of an approved application
package main

import (
	"html/template"
	"net/http"
)

func deployHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/deploy.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	t.Execute(w, data)
}
