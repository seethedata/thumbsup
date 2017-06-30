//applicationList.go lists the status of the applications in the system
package main

import (
	"html/template"
	"log"
	"net/http"
)

func applicationListHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/applicationList.tmpl", "templates/head.tmpl")
	check("Parse template", err)
	var data Payload
	data.URL = mainURL
	data.Applications = getApplications()
	if err != nil {
		log.Fatalf("Failed to create json: %v", err)
	}

	t.Execute(w, data)
}
