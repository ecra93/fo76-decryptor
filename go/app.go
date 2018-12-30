package main

import (
  "html/template"
  "net/http"
  "log"
)

type Input struct {
  Keyword string
  Letters string
  Numbers string
}

type Output struct {
  Phrase string
  Code   string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("./index.html"))
  w.Header().Set("Content-Type", "text/html")

  inputs := Input {
    Keyword: r.FormValue("keyword"),
    Letters: r.FormValue("letters"),
    Numbers: r.FormValue("numbers"),
  }

  _ = inputs

  tmpl.Execute(w, nil)
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()

  tmpl := template.Must(template.ParseFiles("./decrypt.html"))
  w.Header().Set("Content-Type", "text/html")

  input := Input {
    Keyword: r.Form.Get("keyword"),
    Letters: r.Form.Get("letters"),
    Numbers: r.Form.Get("numbers"),
  }

  var keycards []Keycard
  var k *Keycard
  for i, char := range input.Letters {
    k = NewKeycard(byte(char), byte(input.Numbers[i]))
    keycards = append(keycards, *k)
  }

  p, c := Decrypt(input.Keyword, keycards)

  output := Output {
    Phrase: p,
    Code:   c,
  }

  tmpl.Execute(w, output)
}

func main() {

  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/decrypt", decryptHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
