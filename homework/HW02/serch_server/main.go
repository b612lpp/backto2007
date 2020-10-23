package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type dataStruct struct {
	SearchString string   `json:"search"`
	Sites        []string `json:"sites"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", firstHandle)
	router.HandleFunc("/search", handleUserReques)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
func firstHandle(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("Привет, мир!"))
}

//func handles request and returns structure with original search request and sites which contain search word
func handleUserReques(serverResp http.ResponseWriter, incReq *http.Request) {
	var searchReqest dataStruct
	decoder := json.NewDecoder(incReq.Body)
	err := decoder.Decode(&searchReqest)
	if err != nil {
		log.Println(err)
	}
	//in case if specified sites do not contain search word this expression will not work and empty structure will be returned. Need to know how to check if json is empty
	if searchEngine(searchReqest) != nil {
		serverResp.Write(searchEngine(searchReqest))
	} else {
		serverResp.Write([]byte("no matches  found"))
	}

}

//function gets parsed structure from original json reqest and returns anoter structure converted to json
func searchEngine(searchReqestStruct dataStruct) []byte {
	var resultStruct dataStruct
	for _, v := range searchReqestStruct.Sites {

		if strings.Contains(getBody(v), searchReqestStruct.SearchString) {
			resultStruct.Sites = append(resultStruct.Sites, v)
		}

	}
	if resultStruct.Sites != nil {
		resultStruct.SearchString = searchReqestStruct.SearchString
	}
	resultJSON, _ := json.Marshal(resultStruct)
	return resultJSON
}

//function extracts body and returns it as string
func getBody(site string) string {
	req, reqErr := http.Get(site)
	if reqErr != nil {
		log.Println(reqErr)
	}
	defer req.Body.Close()
	body, bodyErr := ioutil.ReadAll(req.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
	}
	bodyStr := string(body)

	return bodyStr
}
