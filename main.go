package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Uint("port", 8080, "proxy listen address")
	user := flag.String("user", "", "proxy auth user")
	pass := flag.String("pass", "", "proxy auth pass")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	auth.ProxyBasic(proxy, "Auth", func(usr, passwd string) bool {
		return usr == *user && passwd == *pass
	})
	proxy.Verbose = true

	log.Printf("Started on port %d\n", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(*port)), proxy))
}
