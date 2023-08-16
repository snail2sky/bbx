package http_echo

import (
	"fmt"
	"github.com/snail2sky/bbx/types"
	"log"
	"net/http"
)

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println(fmt.Sprintf("HTTP echo server handle request: <%s %s %s> from client (%s) ", request.Method, request.RequestURI, request.Proto, request.RemoteAddr))
	err := request.Write(writer)
	if err != nil {
		log.Println(err)
	}
}

func Run(data *types.HTTPEchoData) {
	var listenOn = fmt.Sprintf("%s:%d", data.Host, data.Port)
	log.Println("HTTP echo server listen on:", listenOn)

	http.HandleFunc("/", EchoHandler)
	log.Println(http.ListenAndServe(listenOn, nil))
}
