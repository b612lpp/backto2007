package posts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ConnData contains information to create connection string
type ConnData struct {
	Driver  string `json:"driver"`
	ConnStr string `json:"servaddr"`
}

//GetConnParams asks config service and returns structure with params
func GetConnParams(service string) (conf1 ConnData) {

	requestParams, _ := http.Get("http://127.0.0.1:8081/?servicename=" + service)
	body, _ := ioutil.ReadAll(requestParams.Body)

	_ = json.Unmarshal(body, &conf1)
	defer requestParams.Body.Close()
	//resulString = conf1.Driver + "," + conf1.ConnStr
	return conf1
}
