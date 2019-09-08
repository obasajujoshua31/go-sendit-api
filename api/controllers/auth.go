package controllers

import (
	"fmt"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LoginUser")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Account")
}
