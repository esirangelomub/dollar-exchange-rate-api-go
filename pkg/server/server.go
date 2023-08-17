package server

import (
	"log"
	"net/http"
	"time"
)

func InitializeServer() {
	http.HandleFunc("/cotacao", handleRequests)
	http.ListenAndServe(":8080", nil)
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("Handler started")
	defer log.Println("Handler ended")
	select {
	case <-time.After(5 * time.Second):
		// log imprime no stdout
		log.Println("Handler done")
		// imprime no browser
		w.Write([]byte("Handler done"))
	case <-ctx.Done():
		// log imprime no stdout
		log.Println("Handler cancel by client")
		// imprime no browser
		http.Error(w, "Handler cancel by client", http.StatusRequestTimeout)
	}
}
