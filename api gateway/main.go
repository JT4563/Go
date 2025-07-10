package main

import (
"log"
"net/http"
"net/http/httputil"
"net/url"

"github.com/gorilla/mux"
)
func main(){
r := mux.NewRouter()
r.PathPrefix("/auth/").Handler(reverseproxy("http://localhost:8001"))

r.PathPrefix("/user/").Handler(reverseproxy("http://localhost:8002"))

r.PathPrefix("/payment/").Handler(reverseproxy("http://localhost:8003"))

log.Println("API gateway running on 8080")
log,fatal(http.Listenandserve(":8080",r))
func reverseProxy(target string) http.Handler {
	url, _ := url.Parse(target) //url take target object a bania debe go ke bojhar jonnyo                       // ✅ 1
	proxy := httputil.NewSingleHostReverseProxy(url)   
	  //works for single target
   //ata target url client theke asa rquest forward korbe 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("🔁 Proxying:", r.URL.Path)          // ✅ 3
		proxy.ServeHTTP(w, r)    //atai this send the original cleint request to the mircroservice and then write back the response to the client                        // ✅ 4
	})
}