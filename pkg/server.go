package mocksrv

import (
	"github.com/gin-gonic/gin"
)

func Run(config ConfigRoot, handler *gin.Engine) {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handler)
	// s := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// }
	// log.Fatal(s.ListenAndServe())

	handler.Run(config.Port)
}
