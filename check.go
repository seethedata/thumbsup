//check.go is an error logging function
package main

import (
	"log"
)

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}
