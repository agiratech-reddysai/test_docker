// main.go
package main

import (
  "fmt"
  "path"
  "runtime"
  "net/http"
  "html/template"
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
  _, filename, _, ok := runtime.Caller(0)
  if !ok {
    fmt.Println("No caller information")
  }
  tmpl := template.Must(template.ParseFiles(path.Dir(filename)+"/views/forms.html"))
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