package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	remoteUrl, err := url.Parse(r.Header.Get("X-target"))
	if err != nil {
		w.Write([]byte("500: Invalid url:" + err.Error()))	
		return
	}

	r.URL.Host = remoteUrl.Host
	r.URL.Scheme = remoteUrl.Scheme
	r.Header.Del("X-target")
	r.Host = remoteUrl.Host

	proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
	proxy.ServeHTTP(w, r)
	return
}
