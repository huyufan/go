package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	run := mux.NewRouter()
	go h.run()
	run.HandleFunc("/ws", myws)
	run.HandleFunc("/", local)

	if err := http.ListenAndServe("127.0.0.1:11112", run); err != nil {
		fmt.Println("err:", err)
	}
}
