package controllers

import (
    "fmt"
    "net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Get Method!!")
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Post Method!!")
}

func PutMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Put Method!!")
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Delete Method!!")
}
