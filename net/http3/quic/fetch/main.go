// go get -u github.com/lucas-clemente/quic-go
// go get -u github.com/marten-seemann/qpack
// go run main.go https://www.google.com

package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	url := os.Args[1]

	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	defer roundTripper.Close()

	hclient := &http.Client{
		Transport: roundTripper,
	}

	rsp, err := hclient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
