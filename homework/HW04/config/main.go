package main

import (
	"encoding/json"
	"net/http"
)

type configStr struct {
	Driver   string
	ServAddr string
}

//m emulates db with connection parameters
var m = map[string]configStr{
	"blog": {
		Driver:   "mysql",
		ServAddr: "root:password@tcp(127.0.0.1:3306)/blog"},
	"smthelse": {
		Driver:   "postgres",
		ServAddr: "postgres://username:password@url.com:5432/dbName"},
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", mainResponse)
	http.ListenAndServe("127.0.0.1:8081", handler)
}

//function returns json file which contains connection parameters
func mainResponse(servResp http.ResponseWriter, reqHand *http.Request) {

	requestedConfig := reqHand.URL.Query().Get("servicename")

	config, _ := json.Marshal(m[requestedConfig])

	servResp.Write(config)

}
