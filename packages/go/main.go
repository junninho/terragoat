package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // XSS vulnerability
    name := r.URL.Query().Get("name")
    fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
} 