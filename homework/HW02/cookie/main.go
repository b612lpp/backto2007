package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", setCookie)
	router.HandleFunc("/user", getCookie)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", router))
}
func setCookie(serverResp http.ResponseWriter, incReq *http.Request) {
	cookie := http.Cookie{Name: "UserName", Value: "Uasya"}
	http.SetCookie(serverResp, &cookie)
}

func getCookie(serverResp http.ResponseWriter, incReq *http.Request) {
	extractedCookies := incReq.Cookies()
	if extractedCookies[1].Value != "" {
		fmt.Fprintf(serverResp, "Hi %s", extractedCookies[1].Value) //need more cookies
	} else {
		serverResp.Write([]byte("Visit root page to set cookie"))
	}
}
