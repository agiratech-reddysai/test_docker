// main.go
package main

import (
  "html/template"
  "net/http"
  "fmt"
)

type ContactDetails struct {
  Email   string
  Subject string
  Message string
}

func main() {

  http.HandleFunc("/", formHandler)

  http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("views/forms.html"))
  if r.Method != http.MethodPost {
    tmpl.Execute(w, nil)
    return
  }

  details := ContactDetails{
    Email:   r.FormValue("email"),
    Subject: r.FormValue("subject"),
    Message: r.FormValue("message"),
  }

  // do something with details
  _ = details
  fmt.Println(details)

  tmpl.Execute(w, struct{ Success bool }{true})
}