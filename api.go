package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Created by apsisca_ariadne_alethia")
    })

    http.ListenAndServe(":8080", nil)
}

