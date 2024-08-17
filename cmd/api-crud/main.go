package main

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles(filepath.Join("frontend/templates", "index.html"))
        if err != nil {
            http.Error(w, "Could not load template", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    })

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}