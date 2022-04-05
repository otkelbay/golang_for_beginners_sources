package main

import (
	"log"
	"net/http"
)

var studentsList []string

func main() {
	studentsList = []string{
		"Yerlan", "Arman", "Nurlan",
	}

	http.HandleFunc("/students-list", studentsListHandler)

	err := http.ListenAndServe("localhost:8888", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}

func studentsListHandler(w http.ResponseWriter, r *http.Request) {
	list := ""

	for _, student := range studentsList {
		list += `<li>
		  ` + student + `
		</li>`
	}

	body := `
		<!DOCTYPE html>
		<html>
		<body>
		
		<h2>Students list</h2>
		
		<ul>
		  ` + list + `
		</ul>
		
		</body>
		</html>
	`

	w.Write([]byte(body))
}
