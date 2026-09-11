package main

import (
  "html/template"
  "log"
  "net/http"
  "asciiartwebstylize/core"

)

type pageData struct {
  Result string
  Input string
}

var templ = template.Must(
  template.New("index.html").Funcs(template.FuncMap{
    "safe": func(s string) template.HTML {
      return template.HTML(s)
    },
  }).ParseGlob("./templates/*.html"),
)



func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  http.HandleFunc("/", HomeHandler)
  http.HandleFunc("/ascii-art", GenerateHandler)

  log.Println("Server running on port 4000")
  log.Fatal(http.ListenAndServe(":4000", nil))
}



func HomeHandler(w http.ResponseWriter, r *http.Request) {

  if r.URL.Path != "/"{
	w.WriteHeader(http.StatusNotFound)
	templ.ExecuteTemplate(w,"NotFound.html",pageData{} )
	return
  }

  if err := templ.ExecuteTemplate(w, "index.html", pageData{}); err != nil {
    log.Fatal("Error loading template:", err)
  }
}


func GenerateHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    return
  }


  text := r.FormValue("text")
  banner  := r.FormValue("banner")

  bannerType := map[string]bool{
	"thinkertoy":true,
	"standard":true,
	"shadow":true,
  }

  if text == "" || !bannerType[banner]{
	http.Error(w, "Invalid input", http.StatusBadRequest)
	return
  }

  result := core.Generate(text, banner)

  data := pageData{
    Result: result,
	Input:text,
}

  if err := templ.ExecuteTemplate(w,"index.html", data); err != nil {
	http.Error(w, "Invalid input", http.StatusInternalServerError)
    log.Println("Template error:", err)
  }
}
