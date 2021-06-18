package mocksrv

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func Handler(roots []Root) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		reqId := rand.Int31()
		reqBody, _ := ioutil.ReadAll(req.Body)
		log.Printf("(%8X) Request: [Method: %v, Path: %v, Body: %v]",
			reqId,
			req.Method,
			req.URL.Path,
			string(reqBody))
		for _, r := range roots {
			if r.Path == req.URL.Path && r.Method == req.Method {
				log.Printf("(%8X) Found handler %v", reqId, r.Name)
				headers := r.Response.Headers
				for k, v := range headers {
					resp.Header().Add(k, v)
				}
				if r.Response.Code != 0 {
					resp.WriteHeader(r.Response.Code)
				}
				fmt.Fprintf(resp, r.Response.Body)
				return
			}
		}

		log.Printf("(%8X) No handlers found", reqId)
	}
}
