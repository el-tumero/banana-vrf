package api

import (
	"fmt"
	"net/http"

	"github.com/el-tumero/banana-vrf-client/proposals"
)

func readReqHandler(w http.ResponseWriter, r *http.Request) {
	var out []byte = []byte{}
	for _, p := range proposals.GetStorage() {
		out = append(out, []byte(p.Num.String())...)
		out = append(out, 10) // \n
	}
	w.Write(out)
}

func CreateHttpServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/read", readReqHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return server
}
