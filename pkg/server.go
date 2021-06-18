package mocksrv

import (
	"log"
	"net/http"
)

func Run(handler func(http.ResponseWriter, *http.Request)) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
