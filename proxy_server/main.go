package main

import (
	//	"encoding/json"
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	// log the headers
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			log.Print(r.Header)
			//if auth, ok := r.Header["Authorization"]; ok {
			if _, ok := r.Header["Authorization"]; ok {
				log.Print("Do something with Authorization")
			} else {
				log.Print("Authorization is missing")
				return r, goproxy.NewResponse(r, "application/json", http.StatusBadRequest,
					`{"message": "Bad Request"}`)
			}
			return r, nil
		})
	/*
	   proxy.OnRequest(goproxy.DstHostIs("www.reddit.com")).DoFunc(
	       func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	           h, _, _ := time.Now().Clock()
	           if h >= 8 && h <= 17 {
	               return r, goproxy.NewResponse(r,
	                   goproxy.ContentTypeText, http.StatusForbidden,
	                   "Don't waste your time!")
	           } else {
	               ctx.Warnf("clock: %d, you can waste your time...", h)
	           }
	           return r, nil
	       })
	*/
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	proxy.Verbose = *verbose
	// log the headers
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
