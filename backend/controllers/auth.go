package controllers

import (
	"fmt"
	"net/http"
)


func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST"{
		fmt.Println("t un fdp")
	}
	fmt.Println(r)

}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
