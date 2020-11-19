package main

import (
	"crypto/md5"
	"net/http"
	_ "net/http/pprof"
)

func handle(w http.ResponseWriter, r *http.Request) {
	hash := md5.Sum([]byte(r.RemoteAddr))
	w.Write(hash[:])
}

func main() {
	// runtime.GOMAXPROCS(1)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)

	// ln, err := reuseport.Listen("tcp4", "localhost:8080")
	// if err != nil {
	// 	log.Fatalf("error in reuseport listener: %s", err)
	// }

	// if err := fasthttp.ListenAndServe("localhost:8080", func(ctx *fasthttp.RequestCtx) {
	// 	// hash := md5.Sum(ctx.RemoteIP())
	// 	ctx.WriteString("Hello World!")
	// }); err != nil {
	// 	log.Fatalf("error in fasthttp Server: %s", err)
	// }
}
